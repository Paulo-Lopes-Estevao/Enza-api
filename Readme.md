# Microservice API World

- Country

|Method|Endpoint                        |Description                                                    |BP |QP |
|---   |---                             |---                                                            |---|---|
|GET   |/coutry                 |It returns all country                        |---|---|
|GET   |/coutry/{name}                 |It returns search country                        |---|---|


<br>

<table>
<tr>
<th>Response</th>
<th>On Error</th>
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