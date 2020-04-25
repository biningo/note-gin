# 前言
==我为什么要写一个个人的云笔记？==
(⊙o⊙)…额额额😄。。。这件事儿还得从一只蝙蝠说起......
好了_(:з」∠)_不想扯了......

**~~有以下几个原因：~~**
1. 突然发现有道云笔记不是自己的 我的隐私都在别人的库里😱。并且一直想有一个自己的云笔记。
2. 打开全球最大同性交友平台gay....嗯嗯嗯...错了，Github。发现开源的云笔记软件很少几乎没有
3. 想做一个前后端分离的项目，并且想开源分享给大家。
4. 一只蝙蝠的传说。（大家应该清楚🙃）
---
以上就是我一时冲动，不对。。。一时兴起，所以抄起我的小本本就开干！！来一场说撸就撸的guan......额不对，说撸就撸的码！（时不时就开车。老司机快上车！！🚜）
<br><br>
# 用到的相关技术栈

**前端**
- Vue（包括VueRouter和axios异步请求工具）
- ElementUI
- mavonEditor也就是你们现在看到的markdown编辑器【==码字贼清爽，放大效果更佳！==】

**后端**
- Gin【golang】+Gorm
- Redis+MySql【本项目大量用到了==Redis== 主要用于缓存和消息暂存 **互联网性能利器**，速度快的一逼👍！！！】
- golang相关技术栈（go-redis、gomail、cron定时任务等）

**运维**
- nginx前后端分离部署【前端通过nginx代理跨域】（应该有不少和我一样被跨域折腾的人不人鬼不鬼的小伙伴吧，没事，后面Biningo带你开车.....）
---
不得不说go的部署真的是非常方便！！向改就改，想部署就部署，想停机就停机！！
不过Biningo也是用很多精力部署上去的，主要还是出了各种bug修复一下又修复一下然后又上传这样子。
如果觉得项目有帮助，可以请小弟喝茶。付款码就不贴了，这样显得我Low！！我是个有尊严的码农！
实在要感谢我可以邮箱联系我[biningo@yeah.net](biningo@yeah.net)，打声招呼！【学大学教授 用邮箱沟通，显得我专业😑】

<br><br>
# 注意事项
1、编辑区的文章会自动缓存，可以点击清空即可清空  不然就是更新上次的编辑了

2、编辑器可以放大编辑更方便，同时目录也可也直接在编辑区里面修改

3、没有对文章判断空操作，所以注意，也就是空的文章，没有标题的文章也可以保存，【不过我相信你应该没有那么蛋疼🙄】

4、文章的图片上传是传到七牛云的，然后返回图片链接到markdown，【ps:七牛云也是要钱的，大家试试效果传一张两张就行，毕竟都是流量啊😁】

5、删除是递归删除的，如果你删了目录，那么下面的所有文章都会被删除，而且不可以恢复。如果你仅仅是删除单个文章，目录还在的，那么可以到回收站里面恢复文章，如果回收站里面文章全部清空，那就是永久删除！

6、目录操作也会缓存，方便下次进入你上次操作过的目录

7、新建的文章可以在目录区里面建，也可在编辑区里面先清空上次缓存，再写新的文章，选择目录，会制动创建到相应的目录。

8、文章查看直接点击文章标题即可，退出直接点×或者按**ESC**快捷键即可，本编辑器也可也操作快捷键，具体请查看[https://github.com/hinesboy/mavonEditor/blob/master/README.md](https://github.com/hinesboy/mavonEditor/blob/master/README.md)

9、管理界面支持文章markdown文件下载，批量下载过于麻烦，没有实现，下载的仅仅是markdown文本，方便文章搬家和永久保存。

10、暂时想不到了 ...呃呃呃   👓

---
该叮嘱的还得叮嘱你说是不是？😁
<font color=red>
**如果大家要改造自己用的话也可联系我，我教你部署方法。不超过十分钟就可以部署成功，拥有一个个人云笔记！激不激动？刺不刺激？😁**
</font>
**也希望大家能多做笔记，好记性不如烂键盘嘛😄  同时我也自己部署了一份自用哦**

<br><br>
# 最后想说的话
> 由衷希望每一位coder能完成自己的梦想，希望你们的付出都能得到回报！😏加油！




特别喜欢一位大佬的话，姑且先叫他**码农斯基**吧
> **对内，过好每一天。对外，做好每件事。-----码农斯基**

共勉！

我叫 **biningo**  觉得有用，**点个star，中华传统美德😘 难道你想白嫖我嘛😁**


前端地址：[https://github.com/biningo/note-vue](https://github.com/biningo/note-vue)
后端地址：[https://github.com/biningo/note-gin](https://github.com/biningo/note-gin)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200228170439491.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDU4NDI5Mw==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200228170622303.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDU4NDI5Mw==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200228170455569.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDU4NDI5Mw==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/2020022817050670.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDU4NDI5Mw==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200228170515801.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDU4NDI5Mw==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200228170528381.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDU4NDI5Mw==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200228170547784.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDU4NDI5Mw==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200228170604671.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDU4NDI5Mw==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200228170654877.PNG?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80NDU4NDI5Mw==,size_16,color_FFFFFF,t_70)