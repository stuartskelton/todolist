rm .todos.json
rm tdo
go build -v -o tdo todo-cli.go
./tdo init
./tdo l
./tdo list
echo 'Add'
./tdo add foo
./tdo a foobar
./tdo a foobär
./tdo a foobar 1 @home
./tdo a foobar 2 @away
./tdo a foobar 3 @home +project1
./tdo a foobar 4 @home +project1 due tod
echo 'EDIT'
./tdo e 1 due tom
./tdo l
echo 'COMPLETE'
./tdo complete 4
./tdo complete 2
./tdo c 5
./tdo l
echo 'UNCOMPLETE'
./tdo uncomplete 5
./tdo uc 2
./tdo l
echo 'AC TODOS'
./tdo ac
./tdo l
echo 'AR TODOS'
./tdo ar 2
./tdo l
echo 'UAR TODOS'
./tdo uar 4
./tdo l
echo 'DELETE TODOS'
./tdo d 7
./tdo l
echo 'EXPAND TODOS'
./tdo ex 3 +mooper: adasdasdasdasdasd, aasdasdawerwerwer, wieurowiurpowieur @pppl
./tdo l
./tdo l due tod
echo 'P TODOS'
./tdo p 8
./tdo l
echo 'UP TODOS'
./tdo up 8
./tdo l
# ./tdo gc
./tdo l
echo 'LIST COMPLETED TODOS'
./tdo l completed tod
echo 'RESET THE TODOS'
git checkout .todos.json

