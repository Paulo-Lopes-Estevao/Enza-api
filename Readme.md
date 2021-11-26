# Microservice API Enza

Enza API é um projeto totalmente grátuito e tem como objectivo apresentar informações dos Paises por *Continentes*.

- Desenvolvido em `Golang` ela apresenta nformações como : 
  - Nome do Pais
  - Bandeira
  - Mapa
  - Capital
  - Moeda
  - Idioma
  - População
  - Codigo telefonico

## Endpoints

| Method | Endpoint | Description | BP | QP |
| :---: | :---: | :---: | :---: | :---: |
| GET | /v1/coutry | It returns the details of all countries |
| GET | /v1/coutry/{name} | It returns the details of the referred country |

## Response 

```json
{
    "Flag" : string,  
    "Location" : string,  
    "Keywords" : []string,
    "Name" : string,  
    "Capital" : string,  
    "Population" : float64, 
    "Area" : float64, 
    "Callingcode" : string,  
}
```

## JSON data response

> Example List `/v1/africa` 

```json
{
      "flag": "",
      "location": "",
      "keywords": [""],
      "name": "",
      "capital": "",
      "currency": "",
      "language": "",
      "population": 0,
      "area": 0,
      "callingcode": ""
    }
```

---

&copy; 2021 Paulo-Lopes-Estevao
