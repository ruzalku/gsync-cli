# gsync-cli
google drive sync cli


## Authintication to Google Account

1. Create project on Google Cloud
2. Create Auth Client on _API & Services > OAuth consent screen > Client_ with type "Desktop App" and save JSON file.
3. Add scopes on _API & Services > OAuth consent screen > Data Access_ 
**Scopes:**
```
https://www.googleapis.com/auth/documents
https://www.googleapis.com/auth/drive
https://www.googleapis.com/auth/drive.file

```
4. Install gcloud and run commands:
```bash
gcloud init
gcloud auth application-default login \
  --client-id-file=path to clients JSON file \
  --scopes="https://www.googleapis.com/auth/cloud-platform,https://www.googleapis.com/auth/documents,https://www.googleapis.com/auth/drive"
```

## Building project
1. Move to root of project and run command:
```bash
go build -o gsync-cli
```
2. You will give file gsync-cli and you can use it to run commands


## Examples
```
gsync-cli save /path/to/file.docs
```
