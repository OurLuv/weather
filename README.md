## To launch the application
```bash
make build && make run
```

## Request examples
### Get city list
```bash
http://localhost:8080/city/list
```

### Get short forecast for certain city
```bash
http://localhost:8080/city/5/short-forecast
```

### Get detailed forecast for certain city and time
```bash
http://localhost:8080/city/5?time=2024-07-11 21:00:00
```
