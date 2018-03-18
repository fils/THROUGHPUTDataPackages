Carp Lake Data


## Data

* Data from XYZ
* Also held in Core repo SDF
* Some data present in Neotoma

## People


### Notes from the Source

General: 

* note 1
* note 2


## Citation

Citation info here....

## Automation

The XLSX document were not altered.  They were converted simply to CSV but are
structured in a manner that does not lend itself to automated reformated into a machine
readable format.

The PDF and the XLSX files were processed through the Tika docker file using the 
script located in the script directory.

https://hub.docker.com/r/logicalspark/docker-tikaserver/

docker run -d -p 9998:9998 logicalspark/docker-tikaserver

To extract meta data from a file you can use cURL like this:

Call examples  (no accept style return KV pairs one set per line)
```
curl -H "Accept: application/json" -T testWORD.doc http://localhost:9998/meta
curl -H "Accept: application/json"  -T "Whitlock et al 2000 Env history, tephrostrat Carp Lake.pdf" http://0.0.0.0:9998/meta
curl -T "Whitlock et al 2000 Env history, tephrostrat Carp Lake.pdf" http://localhost:9998/meta
curl -H "Accept: application/json"  -T "Whitlock et al 2000 Env history, tephrostrat Carp Lake.pdf" http://0.0.0.0:9998/tika
curl -T "Whitlock et al 2000 Env history, tephrostrat Carp Lake.pdf" http://localhost:9998/tika
curl -T "Carp Lake log sheets sm.pdf" http://localhost:9998/tika 
curl -X PUT --data-binary @"Whitlock et al 2000 Env history, tephrostrat Carp Lake.pdf" http://localhost:9998/tika --header "Content-type: application/pdf"
curl -X PUT --data-binary @"Carp Lake log sheets sm.pdf" http://localhost:9998/tika --header "Content-type: application/pdf"

```

Example for CSV version  (XLSX via application/vnd.ms-excel   ??)
```
curl -X PUT -H "Content-Disposition: attachment; filename=carpLakeCoreStratigraphy.xlsx" --upload-file carpLakeCoreStratigraphy.xlsx http://localhost:9998/detect/stream

curl -X PUT -H "Content-Disposition: attachment; filename=carpLakeCoreStratigraphy.xlsx" --upload-file carpLakeCoreStratigraphy.xlsx http://localhost:9998/tika

```





## License

License info here...


 