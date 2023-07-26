# Multithreading 

This little API follows the requirements of the task suggested by `FullCycle` in the `GoExpert` course.

The task is:

```
Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

    * https://cdn.apicep.com/file/apicep/" + cep + ".json
 	* ex: https://cdn.apicep.com/file/apicep/88040-050.json
	* result: {
        "code":"88040-050",
        "state":"SC",
        "city":"Florianópolis",
        "district":"Pantanal",
        "address":"Servidão Martinho Leandro dos Santos",
        "status":200,
        "ok":true,
        "statusText":"ok"
        }

    * http://viacep.com.br/ws/" + cep + "/json/
	* ex: http://viacep.com.br/ws/88040-050/json/
	* result: {
        "cep": "88040-050",
        "logradouro": "Servidão Martinho Leandro dos Santos",
        "complemento": "",
        "bairro": "Pantanal",
        "localidade": "Florianópolis",
        "uf": "SC",
        "ibge": "4205407",
        "gia": "",
        "ddd": "48",
        "siafi": "8105"
        }

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.
```

---

## Requirements
* You need to have installed locally:
  * Golang (version > 1.19)

---

## Run the API

1) Clone/download the repository to a local folder;

2) Via terminal(command line), access the cloned repository folder;

3) Run this command to start the API:
    * ``` go run main.go ```

---

## Result

* After running the command, the result of the API will appear on the command line, showing the fastest request made.

* Depending on the number of requests made to the **CDN API CEP**, they can block the requests and return 2 types of response statuses: **429** and **403**. These statuses are "described" in the `internal_test.go` file.