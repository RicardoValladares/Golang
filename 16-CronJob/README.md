
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

si gestionas los servicios con otro sistema, aqui documentacion:
https://www.laboratoriolinux.es/index.php/-noticias-mundo-linux-/software/28866-administrar-servicios-con-systemd-sysv-openrc-y-runit.html
el mas comun es: Sysvinit

