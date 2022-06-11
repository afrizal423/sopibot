#!/bin/bash
SECONDS=0
for i in {1..1200}
do
   ./sopibot
   # echo "Welcome $i times"
done
ELAPSED="Elapsed: $(($SECONDS / 3600))hrs $((($SECONDS / 60) % 60))min $(($SECONDS % 60))sec"
echo "Elapsed: $(($SECONDS / 3600))hrs $((($SECONDS / 60) % 60))min $(($SECONDS % 60))sec"
