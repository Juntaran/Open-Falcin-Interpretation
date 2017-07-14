# Open-Falcon-Interpretation
Interpretion of Open-Falcon


批量修改 import 

```bash
cd $GOPATH/src/falcon-plus
sed -i "s/github.com\/open-falcon\/falcon-plus/falcon-plus/g" `grep github.com/open-falcon/falcon-plus -rl`
```