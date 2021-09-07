# Microservice API World

- ## Country

|Method|Endpoint                        |Description                                                    |BP |QP |
|---   |---                             |---                                                            |---|---|
|GET   |/v1/coutry                 |It returns all country                        |---|---|
|GET   |/v1/coutry/{name}                 |It returns search country                        |---|---|


<br>

<table>
<tr>
<th>Response</th>
</tr>
<tr>
<td>

```json
{
"Flag"       :string,  
"Location   ": string,  
"Keywords   ": []string,
"Name       ": string,  
"Capital    ": string,  
"Population ": float64, 
"Area       ": float64, 
"Callingcode": string,  
}
```
</td>
<td>

</td>
</tr>
</table>

> Examplo List `/v1/africa` 


#
#
#
## json data repository
### 
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