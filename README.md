# study-dockerfile

社の勉強会で使うつもりだけど隠すコードがなにもないのでパブリックにしてます

よろしくなさそうなDockerfileをいい感じにしていく流れで、「Dockerfileがこんな風に変わっていくんやで」というのを見せてやりたいリポジトリだった…

`sudo`を使わずにdockerコマンドが実行できる環境を想定しています

### とりあえず何も考えずに以下を実行してみよう

```bash
git clone git@github.com:nagashima-w/study-dockerfile.git
cd go-hello-huge
docker build -t test:huge .
docker image ls
docker run -d -p 3000:3000
curl localhost:3000
```

たぶん`curl`の実行結果として`Hello World`が返ってきたと思います  
しかしこのDockerfileはイケてないのでいい感じにリファクタリングしてみてください

このリポジトリにあるDockerfileでは、最終的に880MB -> 12.1MBまで軽量化しました
