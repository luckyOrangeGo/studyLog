# 任务：编写用于动态图像构建和多容器测试

目标：这次想象你只是想学习Drupal的管理员和GUI，或者你是一名软件测试人员，并且你需要为Drupal测试一个新的主题。 如果配置正确，这将允许您创建自定义图像，并使用`docker compose up`启动所有内容，包括将重要的数据库和配置数据存储在卷中，以便网站在整个Compose重新启动时记住您的更改。

- 使用您在上次任务（drupal和postgres）中创建的撰写文件作为起点。
- 这次我们将镜像版本从Docker Hub引入。 这样做总是一个好主意，所以一个新的主要版本不会让你感到惊讶。

## Dockerfile

- 首先你需要在这个目录下建立一个自定义的Dockerfile，`FROM drupal:8.2`
- 然后运行apt package manager命令来安装git：`apt-get update && apt-get install -y git`
- 记得在使用`rm -rf /var/lib/apt/lists*`进行apt安装并正确使用`\`和`&&`后进行清理。 你可以在drupal官方图片中找到它们的例子。 在撰写文件下的更多内容。
- 然后改变`WORKDIR /var/www/html/themes`
- 然后使用git克隆主题： `RUN git clone --branch 8.x-3.x --single-branch --depth 1 https://git.drupal.org/project/bootstrap.git`
- 将该行与此行组合在一起，因为我们需要更改文件的权限，并且不希望使用其他图像层来执行此操作（它会让镜像膨胀）。 这个drupal容器以www数据用户身份运行，但构建实际上以root用户身份运行，所以我们经常需要做一些事情，比如`chown`将文件所有者更改为正确的用户：`chown -R www-data:www-data bootstrap`。 记住第一行在末尾需要一个`\`来表示下一行包含在命令中，并且在下一行的开头你应该用`&&`来表示“如果第一个命令成功了，那么也运行这个命令”
- 然后，为了安全起见，在`/var/www/html`处将工作目录改回默认值（从drupal映像）

## 撰写档案

- 我们将在这个用于drupal服务的组合文件中构建一个自定义图像。首先使用之前作业中的Compose文件开始，我们将添加到它，以及更改图像名称。
- 将图像重命名为`custom-drupal`，因为我们想根据官方的`drupal:8.2`制作一个新的图像。
- 我们想通过在`drupal`服务中添加`build:.`来在这个目录下建立默认的Dockerfile。当我们向构建服务添加构建+图像值时，它知道使用图像名称写入到我们的图像缓存中，而不是从Docker Hub中拉出。
- 对于`postgres:9.6`服务，你需要和前面赋值相同的密码，但是也要为`drupal-data:/var/lib/postgresql/data`添加一个卷，这样数据库将在整个Compose重启过程中保持不变。

## 启动容器，配置Drupal

- 像以前一样启动容器，像以前一样配置Drupal web安装。
- 网站出现后，点击顶部栏的“外观”，并注意到一个名为`Bootstrap`的新主题。这是我们用我们的自定义Dockerfile添加的。
- 点击“安装并设置为默认值”。然后点击“返回网站”（左上角），网站界面应该看起来不同。您已经在自己的自定义映像中成功安装并激活了一个新主题，而无需在主机上安装任何其他Docker！
- 如果你退出（ctrl-c），然后`docker-compose`，它将删除容器，但不是卷，所以在下一个`docker-compose up`时，所有东西都会保持原样。
- 要完全清理卷，请在`down`命令中添加`-v`。
