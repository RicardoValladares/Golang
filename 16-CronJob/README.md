
### poner la aplicacion como demonio en systemd linux
```bash
go build -o main.bin
sudo cp main.bin /opt
sudo cp logs.txt /opt
sudo cp CronJob.service /lib/systemd/system
sudo systemctl enable CronJob
sudo systemctl start CronJob 
sudo systemctl status CronJob
journalctl -u CronJob.service
```
