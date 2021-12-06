# blockchain_api

Веб-сервер по предоставлению курса криптавалют, который по обращении к нему в формате:


    ```json
    [
        {
            "symbol":"XLM-EUR"
        }
    ]
    ```
 отдает данные в указанном формате: 


    ```json
    {
        "XLM-EUR": {
            "price": 0.25685,
            "volume": 49644.7076291,
            "last_trade":0.24
        }
    }
    ```

Данные берутся из внешнего источника и сохраняются в БД раз в 30 секунд, а ответ на запрос должен формироваться на основе данных в БД.
- Источник: <https://api.blockchain.com/v3/exchange/tickers>

 - В качестве СУБД используется PostgresQL

# Пример использования:
- git clone https://github.com/SiriusPluge/blockchain_api.git
- sudo docker-compose up

# P.S.: 
- прилагаю коллекцию для Postman (либо)
- 
    ```
    curl --location --request POST 'localhost:8080/api/blockchain' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "symbol":"BCH-USD"
    }'
    ```