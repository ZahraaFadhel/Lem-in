# 🐜 Lem-in Project

## 📝 Overview

Welcome to lem-in Project! This project is a digital version of an ant farm using Go. The goal is to find the quickest path for each ant to go from start to end rooms that are specified in the input text file.
This is the format of the input file : 
```
  number_of_ants
  the_rooms
  the_links
```
A room is defined by "name coord_x coord_y", and will usually look like "Room 1 2", "nameoftheroom 1 6", "4 6 7".

The links are defined by "name1-name2" and will usually look like "1-2", "2-5"

The output will be the input file followed by :
```
Lx-y Lz-w Lr-o ...
```
where x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.

## Usage/Examples

```bash
git clone https://learn.reboot01.com/git/ok/lem-in.git
> cd /lem-in
go run main.go
```
- To run the test file
```
./test.sh
```



## Authors

- [@ok](https://learn.reboot01.com/git/ok)
- [@zfadhel](https://learn.reboot01.com/git/zfadhel)
- [@mohani](https://learn.reboot01.com/git/mohani)

