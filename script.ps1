param([string]$d)
$p = "days/{0:D2}" -f [int]$d
mkdir $p -ErrorAction SilentlyContinue | Out-Null
"desc.md","input.txt","main.go","sample.txt" | ForEach-Object {
    $f = "$p/$_"
    if (-not (Test-Path $f)) { "" | Out-File $f -Encoding utf8 }
}
Write-Host "Day $d ready!" -ForegroundColor Green
