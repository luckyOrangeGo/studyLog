# Linux Shell 系统化学习

## Vagrant Box

>易于配置，可重复使用的环境。

### testbox01指令

#### Box建立

Box = 操作系统映像

```bash
vagrant box add USER/BOX

vagrant box add jasonc/centos7
```

在增加BOX时，出现了ssl错误的情况，而且下载速度极为缓慢，因此我选择使用服务器下载，并使用SecureCRT回传到本地，在本地增加该BOX。

`vagrant box add jasonc/centos7 ./virtualbox.box`

#### Project建立

Vagrant project = Folder with a Vagrantfile.

```bash
mkdir testbox01

cd testbox01
```

现在初始化Vagrant项目:

`vagrant init jasonc/centos7`

`vagrant up`

Vagrant will import the box into VirtualBox and start it.

#### 检查状态

`vagrant status`

#### 进入虚拟机的CMD

The VM is started in headless mode.

`vagrant ssh`

SSH secure shell 是用于连接到Linux系统的网络协议.

#### 备选指令

`vagrant`  ---  列出选项

---
`vagrant halt [VM]` --- 停止 the VM

当你运行这个命令时，你不会失去你在虚拟机上执行的任何工作。虚拟机将仍然存在于VirtualBox中，它将被简单地停止。

可以打开VirtualBox应用程序并验证该虚拟机是否仍然存在，但是已停止。

---
`vagrant up [VM]` ---  开始 the VM，再次启动虚拟机

这次运行命令时，它不需要将虚拟机映像导入到VirtualBox中，因为虚拟机已经存在。

---
`vagrant suspend [VM]`   --- 暂停 the VM

`vagrant resume [VM]`  ---  恢复 the VM

#### 更改名称并重启

Vagrantfile控制虚拟机的行为和设置。

`config.vm.hostname = "testbox01"`

此时，可以运行`vagrant halt`，然后`vagrant up`以激活此更改。但是，Vagrant提供了一个快捷方式,用于重新启动虚拟机，加载新的Vagrantfile配置并再次启动虚拟机。

`vagrant reload`

#### 为虚拟机分配IP地址

创建可以相互通信的虚拟机以及本地工作站。若要给这个虚拟机IP地址“10.9.8.7”。插入Vagrantfile配置

`config.vm.network "private_network", ip: "10.9.8.7"`

重新加载设置文件

`vagrant reload`

通过ping虚拟机来测试连接。

WINDOWS用户，运行命令：`ping 10.9.8.7`

Mac和Linux用户，运行命令：`ping -c 3 10.9.8.7`

#### 参考

vagrant账户的密码是“vagrant”。 root帐户的密码也是“vagrant”。vagrant用户具有完整的sudo（管理）权限，可以进一步配置系统。

#### 退出

`exit`

#### 销毁虚拟机

如果不再需要虚拟机，或者想重新开始虚拟机，运行

`vagrant destroy [VM]`

#### testbox01最终的 Vagrantﬁle

```conf
Vagrant.configure(2) do |config|
    config.vm.box = "jasonc/centos7"
    config.vm.hostname = "testbox01"
    config.vm.network "private_network", ip: "10.9.8.7"
end
```

### 使用多台虚拟机创建一个Vagrant项目

`mkdir multitest && cd multitest`

初始化

`vagrant init jasonc/centos7`

#### multitest项目的 Vagrantfile

```conf
Vagrant.configure("2") do |config|
    config.vm.box = "jasonc/centos7"

    config.vm.define "test1" do |test1|
        test1.vm.hostname = "test1"
        test1.vm.network "private_network", ip: "10.9.8.5"
    end

    config.vm.define "test2" do |test2|
        test2.vm.hostname = "test2"
        test2.vm.network "private_network", ip: "10.9.8.6"
    end
end
```

#### 测试

`vagrant up`

`vagrant status`

连接到test1虚拟机以确认它正在工作，然后退出。

    vagrant ssh test1
    $ exit

连接到test2虚拟机以确认它正在工作。在登录到test2虚拟机时，ping test1虚拟机以证明两台虚拟机可以通过网络相互通信。

    vagrant ssh test2
    ping -c 3 10.9.8.5

当启动虚拟机时，会注意到与此类似的消息

    ==> test2: Mounting shared folders...
    test2: /vagrant => /Users/jason/shellclass/multitest

可以访问驻留在虚拟机内本地计算机上的vagrant proiect目录中的文件。 vagrant项目目录通过`/vagrant`目录进行安装或共享。本地目录中的唯一文件是Vagrantfile。可以从vm中查看文件。在仍然登录test2虚拟机时运行以下命令：

    ls /vagrant
    cat /vagrant/Vagrantfile

`$ exit`

当准备好停止或休息时，停止虚拟机。只要不摧毁虚拟机，总是可以选择挂机。

`vagrant halt`

### Vagrantfile

```conf
Vagrant.configure(2) do |config|
    config.vm.box="jasonc/centos7"

    config.vm.hostname="linuxsvrl"
        config.vm.network "private network", ip:"10.2.3.4"
        config.vm.provider "virtualbox" do |vb|
            vb.gui=true
            vb.memory="1024"
        end

    config.vm.provision "shell", path:"setup.sh"
end
```