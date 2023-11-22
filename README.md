# Usage

```ps1
> $NtpServer = "kr.pool.ntp.org,us.pool.ntp.org,210.72.145.44,pool.ntp.org,kr.pool.ntp.org,us.pool.ntp.org,europe.pool.ntp.org,time1.google.com"

> $NtpServer = $NtpServer.split(",");

> foreach ($i in $NtpServer) {
    $env:SERVER=$i; .\sntp.exe
}
```