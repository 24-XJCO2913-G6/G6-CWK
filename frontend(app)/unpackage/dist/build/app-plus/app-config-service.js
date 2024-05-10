
var isReady=false;var onReadyCallbacks=[];
var isServiceReady=false;var onServiceReadyCallbacks=[];
var __uniConfig = {"pages":["pages/index/index","pages/news/news","pages/upload/upload","pages/msg/msg","pages/my/my","pages/search/search","pages/add-input/add-input","pages/topic-nav/topic-nav","pages/topic-detail/topic-detail","pages/user-list/user-list","pages/user-chat/user-chat","pages/detail/detail","pages/user-set/user-set","pages/user-password/user-password","pages/user-email/user-email","pages/user-userinfo/user-userinfo","pages/user-feedback/user-feedback","pages/about/about","pages/login/login","pages/user-space/user-space","pages/payment/payment","pages/my-post/my-post","pages/user-list/user-list","pages/user-safe/user-safe","pages/user-phone/user-phone","pages/vip/vip","pages/rank/rank","pages/badge/badge","pages/like/like","pages/collect/collect","pages/comment/comment","pages/routes/routes","pages/my-space/my-space"],"window":{"navigationBarTextStyle":"black","navigationBarTitleText":"社区交友","navigationBarBackgroundColor":"#FFFFFF","backgroundColor":"#FFFFFF"},"tabBar":{"color":"#333333","selectedColor":"#1296db","backgroundColor":"#FFFFFF","borderStyle":"black","list":[{"pagePath":"pages/index/index","text":"Home","iconPath":"http://120.46.81.37:8080/app/static/tabbar/index.png","selectedIconPath":"http://120.46.81.37:8080/app/static/tabbar/indexed.png"},{"pagePath":"pages/news/news","text":"Feed","iconPath":"http://120.46.81.37:8080/app/static/tabbar/feed.png","selectedIconPath":"http://120.46.81.37:8080/app/static/tabbar/feed.png"},{"pagePath":"pages/upload/upload","text":"Post","iconPath":"http://120.46.81.37:8080/app/static/tabbar/upload.png","selectedIconPath":"http://120.46.81.37:8080/app/static/tabbar/upload.png"},{"pagePath":"pages/add-input/add-input","text":"Add","iconPath":"http://120.46.81.37:8080/app/static/tabbar/add.png","selectedIconPath":"http://120.46.81.37:8080/app/static/tabbar/add.png"},{"pagePath":"pages/msg/msg","text":"Message","iconPath":"http://120.46.81.37:8080/app/static/tabbar/paper.png","selectedIconPath":"http://120.46.81.37:8080/app/static/tabbar/papered.png"},{"pagePath":"pages/my/my","text":"My","iconPath":"http://120.46.81.37:8080/app/static/tabbar/home.png","selectedIconPath":"http://120.46.81.37:8080/app/static/tabbar/homeed.png"}]},"darkmode":false,"nvueCompiler":"uni-app","nvueStyleCompiler":"weex","renderer":"auto","splashscreen":{"alwaysShowBeforeRender":true,"autoclose":false},"appname":"Been","compilerVersion":"4.07","entryPagePath":"pages/index/index","networkTimeout":{"request":60000,"connectSocket":60000,"uploadFile":60000,"downloadFile":60000}};
var __uniRoutes = [{"path":"/pages/index/index","meta":{"isQuit":true,"isTabBar":true},"window":{"titleNView":{"searchInput":{"align":"center","backgroundColor":"#F5F4F2","borderRadius":"4px","disabled":true,"placeholder":"search","placeholderColor":"#6D6C67"}}}},{"path":"/pages/news/news","meta":{"isQuit":true,"isTabBar":true},"window":{"titleNView":{"searchInput":{"align":"center","backgroundColor":"#F5F4F2","borderRadius":"4px","disabled":true,"placeholder":"search","placeholderColor":"#6D6C67"}}}},{"path":"/pages/upload/upload","meta":{"isQuit":true,"isTabBar":true},"window":{"titleNView":false}},{"path":"/pages/msg/msg","meta":{"isQuit":true,"isTabBar":true},"window":{"navigationBarTitleText":"Messages","enablePullDownRefresh":true,"titleNView":{"buttons":[{"color":"#333333","colorPressed":"#FD597C","float":"left","fontSize":"20px","fontSrc":"http://120.46.81.37:8080/app/static/iconfont.ttf","text":""},{"color":"#333333","colorPressed":"#FD597C","float":"right","fontSize":"20px","fontSrc":"http://120.46.81.37:8080/app/static/iconfont.ttf","text":""}]}}},{"path":"/pages/my/my","meta":{"isQuit":true,"isTabBar":true},"window":{"navigationBarTitleText":"My","titleNView":{"buttons":[{"type":"menu"}]}}},{"path":"/pages/search/search","meta":{},"window":{"enablePullDownRefresh":true,"titleNView":{"searchInput":{"align":"center","backgroundColor":"#F5F4F2","borderRadius":"4px","placeholder":"search","placeholderColor":"#6D6C67"},"buttons":[{"color":"#333333","colorPressed":"#FD597C","float":"right","fontSize":"14px","text":"搜索"}]}}},{"path":"/pages/add-input/add-input","meta":{"isQuit":true,"isTabBar":true},"window":{"titleNView":false}},{"path":"/pages/topic-nav/topic-nav","meta":{},"window":{"navigationBarTitleText":"话题分类"}},{"path":"/pages/topic-detail/topic-detail","meta":{},"window":{"titleNView":{"type":"transparent","buttons":[{"type":"menu"}]}}},{"path":"/pages/user-list/user-list","meta":{},"window":{"titleNView":false}},{"path":"/pages/user-chat/user-chat","meta":{},"window":{"bounce":"none","titleNView":{"buttons":[{"color":"#333333","colorPressed":"#FD597C","float":"right","fontSize":"20px","fontSrc":"http://120.46.81.37:8080/app/static/iconfont.ttf","text":""}]}}},{"path":"/pages/detail/detail","meta":{},"window":{"titleNView":{"buttons":[{"type":"menu","float":"right"}]}}},{"path":"/pages/user-set/user-set","meta":{},"window":{"navigationBarTitleText":"Setting"}},{"path":"/pages/user-password/user-password","meta":{},"window":{"navigationBarTitleText":"Change Passsword"}},{"path":"/pages/user-email/user-email","meta":{},"window":{"navigationBarTitleText":"Bind Email"}},{"path":"/pages/user-userinfo/user-userinfo","meta":{},"window":{"navigationBarTitleText":"Edit Profile"}},{"path":"/pages/user-feedback/user-feedback","meta":{},"window":{"navigationBarTitleText":"意见反馈"}},{"path":"/pages/about/about","meta":{},"window":{"navigationBarTitleText":"关于社区"}},{"path":"/pages/login/login","meta":{},"window":{"titleNView":false}},{"path":"/pages/user-space/user-space","meta":{},"window":{"navigationBarTitleText":"User Space","titleNView":{"buttons":[{"type":"menu"}]}}},{"path":"/pages/payment/payment","meta":{},"window":{"navigationBarTitleText":"Payment","titleNView":{"buttons":[{"type":"menu"}]}}},{"path":"/pages/my-post/my-post","meta":{},"window":{"navigationBarTitleText":"My Post","titleNView":{"buttons":[{"type":"menu"}]}}},{"path":"/pages/user-safe/user-safe","meta":{},"window":{"navigationBarTitleText":"账号与安全"}},{"path":"/pages/user-phone/user-phone","meta":{},"window":{"navigationBarTitleText":"绑定手机号"}},{"path":"/pages/vip/vip","meta":{},"window":{"navigationBarTitleText":"VIP","titleNView":{"buttons":[{"type":"menu"}]}}},{"path":"/pages/rank/rank","meta":{},"window":{"navigationBarTitleText":"Rank","titleNView":{"buttons":[{"type":"menu"}]}}},{"path":"/pages/badge/badge","meta":{},"window":{"navigationBarTitleText":"Medal wall","titleNView":{"buttons":[{"type":"menu"}]}}},{"path":"/pages/like/like","meta":{},"window":{"navigationBarTitleText":"Likes from Others","titleNView":{"buttons":[{"type":"menu"}]}}},{"path":"/pages/collect/collect","meta":{},"window":{"navigationBarTitleText":"Collect from Others","titleNView":{"buttons":[{"type":"menu"}]}}},{"path":"/pages/comment/comment","meta":{},"window":{"navigationBarTitleText":"Comments from Others","titleNView":{"buttons":[{"type":"menu"}]}}},{"path":"/pages/routes/routes","meta":{},"window":{"navigationBarTitleText":"My routes","titleNView":{"buttons":[{"type":"menu"}]}}},{"path":"/pages/my-space/my-space","meta":{},"window":{"navigationBarTitleText":"","enablePullDownRefresh":false}}];
__uniConfig.onReady=function(callback){if(__uniConfig.ready){callback()}else{onReadyCallbacks.push(callback)}};Object.defineProperty(__uniConfig,"ready",{get:function(){return isReady},set:function(val){isReady=val;if(!isReady){return}const callbacks=onReadyCallbacks.slice(0);onReadyCallbacks.length=0;callbacks.forEach(function(callback){callback()})}});
__uniConfig.onServiceReady=function(callback){if(__uniConfig.serviceReady){callback()}else{onServiceReadyCallbacks.push(callback)}};Object.defineProperty(__uniConfig,"serviceReady",{get:function(){return isServiceReady},set:function(val){isServiceReady=val;if(!isServiceReady){return}const callbacks=onServiceReadyCallbacks.slice(0);onServiceReadyCallbacks.length=0;callbacks.forEach(function(callback){callback()})}});
service.register("uni-app-config",{create(a,b,c){if(!__uniConfig.viewport){var d=b.weex.config.env.scale,e=b.weex.config.env.deviceWidth,f=Math.ceil(e/d);Object.assign(__uniConfig,{viewport:f,defaultFontSize:Math.round(f/20)})}return{instance:{__uniConfig:__uniConfig,__uniRoutes:__uniRoutes,global:void 0,window:void 0,document:void 0,frames:void 0,self:void 0,location:void 0,navigator:void 0,localStorage:void 0,history:void 0,Caches:void 0,screen:void 0,alert:void 0,confirm:void 0,prompt:void 0,fetch:void 0,XMLHttpRequest:void 0,WebSocket:void 0,webkit:void 0,print:void 0}}}});
