FROM drupal:8.2

RUN apt-get update && apt-get install -y git \
    && rm -rf /var/lib/apt/lists/*
# 安装git 并移除apt-get lists，可节约10MB空间

WORKDIR /var/www/html/themes

RUN git clone --branch 8.x-3.x --single-branch --depth 1 https://git.drupal.org/project/bootstrap.git \
    && chown -R www-data:www-data bootstrap
# 克隆最新版的bootstrap，单分支，无历史数据，节约时间空间
# 设置最高权限

WORKDIR /var/www/html
