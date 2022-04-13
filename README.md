## Feira Api

### Pré-requisitos
`Go >= 1.18`
`Docker`
`grep`
`make`

### Como usar
Com Docker:
1. Para executar: ````sh make docker/build && make docker/run````

Local:
1. Execute: ```sh go run main.go```

Utilitários:
1. Para executar os testes: ````sh make test````
2. Para testas as condições de corridas: ```sh make test/race```
3. Para ver o coverage dos teste: ```sh make cover/text```
4. Para limpar o ambiente docker: ```sh make docker/clean```
5. Para criar os mocks: ```sh make mocks```

### Sobre as decisões
O projeto foi desenvolvido na linguagem Go, o motivo da escolha foi o conhecimento técnico da linguagem.
O comando `grep` é utilizado pois na geração do build checa-se com o `ldd` o estático gerado.

### Arquitetura
O projeto foi organizado utilizando um simplificação da arquiteturas em camadas. 

`main.go`: inicia o programa chamando o cmd.

`cmd:` é o pacote responsável pela interface de carregamento das variaveis de ambiente e por fazer o setup da aplicação.

`market:` é o pacote que inclui toda a logica de http, banco de dados e casos de uso.

`domain:` contém o modelo de utilização de dominio.

`Dockerfile`: utliliza um builder para gerar o binário go do projeto. E a imagem de execução é apenas uma alpine.

### Exemplo de requests
Home: 
```
curl --request GET \
  --url http://localhost:8888/
```

GetByRegister 
```
curl --request GET \
  --url http://localhost:8888/market/4041-0
```

GetByName: 
```
curl --request GET \
  --url 'http://localhost:8888/market?nome=VILA%20FORMOSA'
```

DeleteByRegister: 
```
curl --request DELETE \
  --url http://localhost:8888/market/4041-0
```

UpdateByRegister: 
```
curl --request PUT \
  --url http://localhost:8888/market/4041-0 \
  --header 'Content-Type: application/json' \
  --data '{
    "LONG": -46550164,
    "LAT": -23558733,
    "SETCENS": 355030885000091,
    "AREAP": 3550308005040,
    "CODDIST": 87,
    "DISTRITO": "VILA FORMOSA",
    "CODSUBPREF": 26,
    "SUBPREFE": "ARICANDUVA-FORMOSA-CARRAO",
    "REGIAO5": "Leste",
    "REGIAO8": "Leste 1",
    "NOME_FEIRA": "VILA FORMOSA",
    "LOGRADOURO": "RUA MARAGOJIPE",
    "NUMERO": "S/N",
    "BAIRRO": "VL FORMOSA",
    "REFERENCIA": "TV RUA PRETORIA"
}'
```

StoreMarket: 
```
curl --request POST \
  --url http://localhost:8888/market \
  --header 'Content-Type: application/json' \
  --data '{
    "LONG": -46550164,
    "LAT": -23558733,
    "SETCENS": 355030885000091,
    "AREAP": 3550308005040,
    "CODDIST": 87,
    "DISTRITO": "VILA FORMOSA",
    "CODSUBPREF": 26,
    "SUBPREFE": "ARICANDUVA-FORMOSA-CARRAO",
    "REGIAO5": "Leste",
    "REGIAO8": "Leste 1",
    "NOME_FEIRA": "VILA FORMOSA",
    "REGISTRO": "4041-0",
    "LOGRADOURO": "RUA MARAGOJIPE",
    "NUMERO": "S/N",
    "BAIRRO": "VL FORMOSA",
    "REFERENCIA": "TV RUA PRETORIA"
}'
```
