$currentPath = Get-Location
Set-Location ..
Invoke-Expression -Command "go build main.go"
$exePath = Get-Location
$fullExePath = Join-Path -Path $exePath -ChildPath "main.exe"
Write-Host $fullExePath
schtasks /create /tn "ObsidianWorker" /tr $fullExePath /sc daily /st 09:00
Set-Location $currentPath