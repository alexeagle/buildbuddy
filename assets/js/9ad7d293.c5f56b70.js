"use strict";(self.webpackChunkbuildbuddy_docs_website=self.webpackChunkbuildbuddy_docs_website||[]).push([[4602],{4137:function(e,t,n){n.d(t,{Zo:function(){return d},kt:function(){return m}});var o=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function r(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,o)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?r(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):r(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,o,i=function(e,t){if(null==e)return{};var n,o,i={},r=Object.keys(e);for(o=0;o<r.length;o++)n=r[o],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(o=0;o<r.length;o++)n=r[o],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var u=o.createContext({}),s=function(e){var t=o.useContext(u),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},d=function(e){var t=s(e.components);return o.createElement(u.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return o.createElement(o.Fragment,{},t)}},p=o.forwardRef((function(e,t){var n=e.components,i=e.mdxType,r=e.originalType,u=e.parentName,d=l(e,["components","mdxType","originalType","parentName"]),p=s(n),m=i,b=p["".concat(u,".").concat(m)]||p[m]||c[m]||r;return n?o.createElement(b,a(a({ref:t},d),{},{components:n})):o.createElement(b,a({ref:t},d))}));function m(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var r=n.length,a=new Array(r);a[0]=p;var l={};for(var u in t)hasOwnProperty.call(t,u)&&(l[u]=t[u]);l.originalType=e,l.mdxType="string"==typeof e?e:i,a[1]=l;for(var s=2;s<r;s++)a[s]=n[s];return o.createElement.apply(null,a)}return o.createElement.apply(null,n)}p.displayName="MDXCreateElement"},9643:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return l},contentTitle:function(){return u},metadata:function(){return s},assets:function(){return d},toc:function(){return c},default:function(){return m}});var o=n(7462),i=n(3366),r=(n(7294),n(4137)),a=["components"],l={slug:"buildbuddy-v1-3-0-release-notes",title:"BuildBuddy v1.3.0 Release Notes",author:"Siggi Simonarson",author_title:"Co-founder @ BuildBuddy",date:"2020-09-30:12:00:00",author_url:"https://www.linkedin.com/in/siggisim/",author_image_url:"https://avatars.githubusercontent.com/u/1704556?v=4",tags:["product","release-notes"]},u=void 0,s={permalink:"/blog/buildbuddy-v1-3-0-release-notes",editUrl:"https://github.com/buildbuddy-io/buildbuddy/edit/master/website/blog/buildbuddy-v1-3-0-release-notes.md",source:"@site/blog/buildbuddy-v1-3-0-release-notes.md",title:"BuildBuddy v1.3.0 Release Notes",description:"We're excited to share that v1.3.0 of BuildBuddy is live on both Cloud Hosted BuildBuddy and open-source via Github and Docker!",date:"2020-09-30T12:00:00.000Z",formattedDate:"September 30, 2020",tags:[{label:"product",permalink:"/blog/tags/product"},{label:"release-notes",permalink:"/blog/tags/release-notes"}],readingTime:3.135,truncated:!1,authors:[{name:"Siggi Simonarson",title:"Co-founder @ BuildBuddy",url:"https://www.linkedin.com/in/siggisim/",imageURL:"https://avatars.githubusercontent.com/u/1704556?v=4"}],prevItem:{title:"Welcoming George Li, Head of Sales",permalink:"/blog/welcoming-george-li-head-of-sales"},nextItem:{title:"BuildBuddy v1.2.1 Release Notes",permalink:"/blog/buildbuddy-v1-2-1-release-notes"}},d={authorsImageUrls:[void 0]},c=[{value:"New to Open Source BuildBuddy",id:"new-to-open-source-buildbuddy",children:[],level:2},{value:"New to Cloud &amp; Enterprise BuildBuddy",id:"new-to-cloud--enterprise-buildbuddy",children:[],level:2}],p={toc:c};function m(e){var t=e.components,l=(0,i.Z)(e,a);return(0,r.kt)("wrapper",(0,o.Z)({},p,l,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("p",null,"We're excited to share that v1.3.0 of BuildBuddy is live on both",(0,r.kt)("a",{parentName:"p",href:"https://app.buildbuddy.io/"}," Cloud Hosted BuildBuddy")," and open-source via",(0,r.kt)("a",{parentName:"p",href:"https://github.com/buildbuddy-io/buildbuddy"}," Github")," and",(0,r.kt)("a",{parentName:"p",href:"https://github.com/buildbuddy-io/buildbuddy/blob/master/docs/on-prem.md#docker-image"}," Docker"),"!"),(0,r.kt)("p",null,"Thanks to everyone using open source and cloud-hosted BuildBuddy. We\u2019ve made lots of improvements in this release based on your feedback."),(0,r.kt)("p",null,"Our focus for this release was on giving users new tools to improve build performance, debug cache hits, and a completely redesigned Cloud & Enterprise experience."),(0,r.kt)("h2",{id:"new-to-open-source-buildbuddy"},"New to Open Source BuildBuddy"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("strong",{parentName:"li"},"Timing profile explorer "),"- Bazel's timing profile is the best way to get detailed insights into where to spend your time optimizing your build. Unfortunately, extracting useful information out of the thousands of events can be challenging without using something like Chrome's profiling tools. Now we've built these tools right into the BuildBuddy timing tab so you can explore this info for any build. See which actions dominate your build's critical path or find out how much time is spent downloading outputs - now with a single click.")),(0,r.kt)("p",null,(0,r.kt)("img",{src:n(7684).Z})),(0,r.kt)("p",null,"Dive into the timing for every action, critical path information, and more."),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("strong",{parentName:"li"},"Cache stats "),"- one of the feature requests we get most frequently is for more information on cache hits and misses. This data can be tricky to get a hold of because it's not made easily available by Bazel's build event protocol. That's why we've introduced BuildBuddy's new cache tab. It gives you a view into cache hits, misses, and writes for every invocation that uses BuildBuddy's gRPC cache. It breaks these numbers down by action cache (AC) and content addressable store (CAS). BuildBuddy also tracks the volume and throughput of cache requests so you can see how much data is moving in and out of the cache - and at what speed.")),(0,r.kt)("p",null,(0,r.kt)("img",{src:n(1343).Z})),(0,r.kt)("p",null,"Get a view into cache performance for every invocation."),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("strong",{parentName:"li"},"Environment variable redaction controls")," - when debugging cache hits, it can be useful to get a full picture of the inputs that are affecting a particular build - like the PATH environment variable. By default, BuildBuddy redacts nearly all environment variables passed into Bazel. We've added controls per invocation that allow you to optionally allow environment variables of your choice to skip redaction. Information on configuring this can be found in our",(0,r.kt)("a",{parentName:"li",href:"https://www.buildbuddy.io/docs/guide-metadata#environment-variable-redacting"}," build metadata docs"),".")),(0,r.kt)("h2",{id:"new-to-cloud--enterprise-buildbuddy"},"New to Cloud & Enterprise BuildBuddy"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("strong",{parentName:"li"},"Redesigned navigation "),"- as BuildBuddy has grown from a debugging tool to a fully-featured platform to debug, analyze, monitor, and share builds across your organization, we've outgrown the minimal navigation setup that has gotten us this far. In Cloud and Enterprise BuildBuddy, we've replaced the top menu bar with a more fully-featured left-nav. This gives us room to add new features like Trends and provides easier access to critical pages like Setup & Docs.")),(0,r.kt)("p",null,(0,r.kt)("img",{src:n(6366).Z})),(0,r.kt)("p",null,"The new navigation makes room for new features."),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("strong",{parentName:"li"},"Trends "),"- BuildBuddy has lots of information about every invocation that it collects. Now with Trends, you can follow how your builds are changing over time. Have all of the cache improvements you've been working on decreased average build time over the last month? Has the addition of a new external dependency significantly increased the length of your slowest builds? Answering these questions is easy with BuildBuddy Trends.")),(0,r.kt)("p",null,(0,r.kt)("img",{src:n(1076).Z})),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("strong",{parentName:"li"},"Redis Pub/Sub support "),"- we've added support for Redis Pub/Sub to significantly improve remote build execution performance. It's completely optional for on-prem deployments, but in our testing it's improved performance for builds with lots of small actions by a factor of 2x. No change is required for Cloud users - just enjoy the faster builds!")),(0,r.kt)("p",null,"That\u2019s it for this release. Stay tuned for more updates coming soon!"),(0,r.kt)("p",null,"As always, we love your feedback - join our",(0,r.kt)("a",{parentName:"p",href:"https://slack.buildbuddy.io"}," Slack channel")," or email us at ",(0,r.kt)("a",{parentName:"p",href:"mailto:hello@buildbuddy.io"},"hello@buildbuddy.io")," with any questions, comments, or thoughts."))}m.isMDXComponent=!0},1343:function(e,t,n){t.Z=n.p+"assets/images/cache-stats-abd2733109aa2543cea9b5155963f142.png"},6366:function(e,t,n){t.Z=n.p+"assets/images/navigation-0fae6ed61c5fcc0732d9647a1244e691.png"},7684:function(e,t,n){t.Z=n.p+"assets/images/timing-profile-e9cc7790c7c7b6b9e27e9af365df69da.png"},1076:function(e,t,n){t.Z=n.p+"assets/images/trends-v0-6e404ed372e31e6c10dd6f9cbb1b39dd.png"}}]);