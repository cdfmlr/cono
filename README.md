# cono

> The next generation CoursesNotifier
> 
> 某种野生的（可能是）微服务的课程提醒系统。

这个项目（好像）已经基本实现了。好长时间没碰了，具体我忘了。。。
也许是可以用了，我好像还在服务器上部署过了（至少是部署了部分进行测试），
那台服务器刚好昨天到期释放了😂，所以具体无从考证。

下面是我的一点模糊记忆：

- 项目（已实现的部分）分成两个微服务：
   - student: 学生（用户）账号服务："学生"增删改查：暴露 gRPC 服务 
   - course: 课程服务："课程"增删改查、上课提醒：提供微信公众号服务
- 这个项目大体上构架的思路是参考 go-kit 的。

注：这个项目的时候我还没真正开始学微服务（现在依然没学多少），这个只是我自以为的微服务。

具体这里记录了一些开发的过程： [conoing.md](conoing.md)。

这个项目给我一个教训：
课程提醒系统其实非常简单了，粗暴的简单服务就足够了，完全没有必要上微服务。

