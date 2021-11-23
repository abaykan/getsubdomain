# getsubdomain
Get subdomain list and check whether they are active or not by each response code. Using API by [c99.nl](https://c99.nl/)

### Installation
```
▶ go install github.com/abaykan/getsubdomain@latest
```
Put your `API_KEY` in `~/.config/c99.txt`:
```
▶ echo "<YOUR_API_KEY>" > ~/.config/c99.txt
```

### Usage
Usage example:
```
▶ echo "kustirama.id" | getsubdomain
```
or
```
▶ cat domainlist.txt | getsubdomain
```

### Example Output
- Single Domain
<img src="https://kustirama.id/images/getsubdomain-single-domain.gif">

- From Domain List
<img src="https://kustirama.id/images/getsubdomain-domain-list.gif">

### Disclaimer
For my own learning purpose. These codes are messy af. Feel free to contribute so I know how to code properly.
