#!/bin/bash

echo "You died."

echo "Do you like coffee bro? (y/n)"
read coffee

if [[ $coffee == "y" ]]; then
  echo "CoffeE 4 Lyf3!"
else
  echo "Do you even code bro?"
fi
