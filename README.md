# Imersão Fullcycle 21 - Home broker

![Imersão Full Stack && Full Cycle](https://events-fullcycle.s3.amazonaws.com/events-fullcycle/static/site/img/grupo_4417.png)

Participe gratuitamente: https://imersao.fullcycle.com.br/

## Requerimentos

Cada projeto tem seus próprios requerimentos, mas uma ferramenta é comum a todos: o Docker.

### Docker

Dependendo do seu sistema operacional, você tem 2 opções para instalar o Docker:

- [Docker Desktop] - Interface gráfica para gerenciar e usar o Docker.
- [Docker Engine] - Apenas a engine do Docker, sem interface gráfica, chamado de Docker Nativo.

Se você tem 8GB ou menos de memória RAM, recomendamos o uso do Docker Engine, pois a interface gráfica do Docker Desktop + a execução dos containers pode consumir praticamente a memória da máquina, caso contrário usar o Docker Desktop é mais prático.

Se você quiser saber mais detalhes sobre isto, veja nosso vídeo [https://www.youtube.com/watch?v=99dCerRKO6s](https://www.youtube.com/watch?v=99dCerRKO6s).

Se você estiver no Windows, use o WSL 2. Veja nosso tutorial [https://github.com/codeedu/wsl2-docker-quickstart](https://github.com/codeedu/wsl2-docker-quickstart).

## Rodar a aplicação

Para rodar os projetos entre em cada pasta e siga as intruções.

Existe uma configuração para rodar todos os projetos em containers Docker, veja o branch **all-in-docker**.

Entre em cada **README.md** dos projetos para rodar os projetos. A ordem de execução é:

2. [Golang](./go/README.md)
1. [Nest.js](./nestjs-api/README.md)
3. [Next.js](./next-frontend/README.md)

## Links e material adicional

* Como montar o melhor ambiente Dev no Windows, Linux e Mac com WSL [https://www.youtube.com/watch?v=O33trWxqVC4](https://www.youtube.com/watch?v=O33trWxqVC4)
* Instalação do Node.js [https://nodejs.org/](https://nodejs.org/)
* Docker [https://www.docker.com/](https://www.docker.com/)
* Tutorial do WSL + Docker [https://github.com/codeedu/wsl2-docker-quickstart]
* Minhas configurações do VSCode [https://github.com/argentinaluiz/my-vscode-settings](https://github.com/argentinaluiz/my-vscode-settings)

## Arquitetura do projeto

![alt text](./arquitetura_projeto.png)
