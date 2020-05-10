#!/usr/bin/env bash


http http://localhost:9876/products ID=1234 Name="Mario Kart"
http http://localhost:9876/products ID=1235 Name="Street Fighter 2 25th Anniversary"
http http://localhost:9876/products ID=1236 Name="Return of the ninja remastered"



http http://localhost:9876/products/1234

http http://localhost:9876/products?name=Fighter