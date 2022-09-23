echo "Rotating..."
mv ./worker.log ./worker.log-old
killall -USR1 app