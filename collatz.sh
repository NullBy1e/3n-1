#!/bin/bash 

collatz () {
  n=$1
  echo $n
  if [[ "$n" -eq 1 ]] || [[ "$n" -eq 2 ]]; then
    echo -e "\n"
  else
    if [[ $(($n % 2)) -eq 0 ]]; then
      final=$(($n/2))
      collatz $final
    else
      final=$((n*3+1))
    fi 
  fi 
}

numbers=$1
for (( i=1; i<=$numbers; i++)); do
  collatz $i 
done
  
