# qcsv

```go
w := NewWriter(writer)
if err := w.WriteAll([][]string{
  "ID", "Name", "Age", "Note",
  "01", "Tom", "20", "My "Best" Friend",
  "02", "Alice", "21", "",
}); err != nil {
    panic(err)
}
w.Flush()
```

## using stdlib csv writer
```csv
ID,Name,Age,Note
01,Tom,20,My ""Best"" Friend,
02,Alice,21,,
```

### using qsv writer
```csv
"ID","Name","Age","Note"
"01","Tom","20","My ""Best""" Friend",
"02","Alice","21","",
```
