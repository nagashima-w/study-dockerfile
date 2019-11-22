# study-dockerfile

社の勉強会で使うつもりだけど隠すコードがなにもないのでパブリックにしてます

よろしくなさそうなDockerfileをいい感じにしていく流れで、「Dockerfileがこんな風に変わっていくんやで」というのを見せてやりたいリポジトリだった…

`sudo`を使わずにdockerコマンドが実行できる環境を想定しています

### とりあえず何も考えずに以下を実行してみよう

```bash
git clone git@github.com:nagashima-w/study-dockerfile.git
cd go-hello-bad
docker build -t study:go-hello-bad .
docker image ls
docker run -d -p 3000:3000 study:go-hello-bad
curl localhost:3000
```

