
### poner la aplicacion como servicio en systemd linux
```bash
go build -o main.bin
sudo cp main.bin /opt
sudo cp CronJob.service /lib/systemd/system
sudo systemctl enable CronJob
sudo systemctl start CronJob 
sudo systemctl status CronJob
```
