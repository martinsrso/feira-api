package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/martinsrso/feira-api/domain"
	_marketHttp "github.com/martinsrso/feira-api/market/delivery/http"
	_marketRepo "github.com/martinsrso/feira-api/market/repository/postgres"
	_marketUsecase "github.com/martinsrso/feira-api/market/usecase"
)

var (
	configFile    string
	dbImportFiles string
	config        *viper.Viper
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api-feira",
	Short: "api-feira command-line interface",
	Long:  "api-feira api to market at SP",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// Execute is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config/default.yaml", "config file (default is $feira-api/config/default.yaml)")
	rootCmd.PersistentFlags().StringVar(&dbImportFiles, "db", "db/DEINFO_AB_FEIRASLIVRES_2014.csv", "db import file (default is $feira-api/db/DEINFO.csv)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config = viper.New()
	if configFile != "" { // enable ability to specify config file via flag
		config.SetConfigFile(configFile)
	}
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()

	// If a config file is found, read it in.
	if err := config.ReadInConfig(); err != nil {
		log.Panicf("config file %s failed to load: %s.\n", configFile, err.Error())
	}
}

func run() {
	initConfig()

	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	file, err := os.Create("db.log")
	if err != nil {
		// Handle error
		panic(err)
	}

	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	dbHost := config.GetString("db_host")
	dbUser := config.GetString("db_user")
	dbPassword := config.GetString("db_pass")
	dbPort := config.GetString("db_port")
	dbName := config.GetString("db_name")
	timezone := config.GetString("timezone")
	port := fmt.Sprintf(":%s", config.GetString("port"))
	importDB := config.GetBool("import_db")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbHost, dbUser, dbPassword, dbName, dbPort, timezone)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	db.AutoMigrate(&domain.Market{})

	if importDB {
		if err := importDBFile(db); err != nil {
			db.Logger.Error(db.Statement.Context, "error on import file")
		}
	}

	r := gin.Default()
	mr := _marketRepo.NewPostgresMarketRepository(db)
	mus := _marketUsecase.NewMarketUsecase(mr, 5*time.Second)

	_marketHttp.NewMarketHandler(r, mus)
	r.Run(port)
}

func importDBFile(db *gorm.DB) error {
	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		return err
	}

	file, err := os.Open(dbImportFiles)
	if err != nil {
		return err
	}
	defer file.Close()

	var entries []domain.Market
	err = gocsv.Unmarshal(file, &entries)
	if err != nil {
		return err
	}

	result := db.Create(entries)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
