"use strict";(self.webpackChunkbuildbuddy_docs_website=self.webpackChunkbuildbuddy_docs_website||[]).push([[7470],{30:function(e,t,a){a.r(t),a.d(t,{default:function(){return d}});var n=a(7294),r=a(9962),i=a(6251),l=a(8890),m=a(3699),o=a(7325);var s=function(e){var t=e.metadata,a=t.previousPage,r=t.nextPage;return n.createElement("nav",{className:"pagination-nav","aria-label":(0,o.I)({id:"theme.blog.paginator.navAriaLabel",message:"Blog list page navigation",description:"The ARIA label for the blog pagination"})},n.createElement("div",{className:"pagination-nav__item"},a&&n.createElement(m.Z,{className:"pagination-nav__link",to:a},n.createElement("div",{className:"pagination-nav__label"},"\xab"," ",n.createElement(o.Z,{id:"theme.blog.paginator.newerEntries",description:"The label used to navigate to the newer blog posts page (previous page)"},"Newer Entries")))),n.createElement("div",{className:"pagination-nav__item pagination-nav__item--next"},r&&n.createElement(m.Z,{className:"pagination-nav__link",to:r},n.createElement("div",{className:"pagination-nav__label"},n.createElement(o.Z,{id:"theme.blog.paginator.olderEntries",description:"The label used to navigate to the older blog posts page (next page)"},"Older Entries")," ","\xbb"))))},c=a(8817);var d=function(e){var t=e.metadata,a=e.items,m=e.sidebar,o=(0,r.Z)().siteConfig.title,d=t.blogDescription,g=t.blogTitle,h="/"===t.permalink?o:g;return n.createElement(i.Z,{title:h,description:d,wrapperClassName:"blog-wrapper"},n.createElement("div",{className:"container margin-vert--lg"},n.createElement("div",{className:"row"},n.createElement("div",{className:"col col--1"}),n.createElement("main",{className:"col col--8"},a.map((function(e){var t=e.content;return n.createElement(l.Z,{key:t.metadata.permalink,frontMatter:t.frontMatter,metadata:t.metadata,truncated:t.metadata.truncated},n.createElement(t,null))})),n.createElement(s,{metadata:t})),n.createElement("div",{className:"col col--3"},n.createElement(c.Z,{sidebar:m})))))}},8890:function(e,t,a){a.d(t,{Z:function(){return g}});var n=a(7294),r=a(6010),i=a(4137),l=a(7325),m=a(3699),o=a(3002),s=a(4175),c={blogPostTitle:"blogPostTitle_3kJS",blogPostDate:"blogPostDate_247f",blogAuthorTitle:"blogAuthorTitle_-PT7",subtitle:"subtitle_2zQw",heading:"heading_1SDJ",headingPhoto:"headingPhoto_1iKi",avatar:"avatar_1f5j",authorName:"authorName_3Kcr",tags:"tags_2aeM",tag:"tag_3Sf8"},d=[(0,l.I)({id:"theme.common.month.january",description:"January month translation",message:"January"}),(0,l.I)({id:"theme.common.month.february",description:"February month translation",message:"February"}),(0,l.I)({id:"theme.common.month.march",description:"March month translation",message:"March"}),(0,l.I)({id:"theme.common.month.april",description:"April month translation",message:"April"}),(0,l.I)({id:"theme.common.month.may",description:"May month translation",message:"May"}),(0,l.I)({id:"theme.common.month.june",description:"June month translation",message:"June"}),(0,l.I)({id:"theme.common.month.july",description:"July month translation",message:"July"}),(0,l.I)({id:"theme.common.month.august",description:"August month translation",message:"August"}),(0,l.I)({id:"theme.common.month.september",description:"September month translation",message:"September"}),(0,l.I)({id:"theme.common.month.october",description:"October month translation",message:"October"}),(0,l.I)({id:"theme.common.month.november",description:"November month translation",message:"November"}),(0,l.I)({id:"theme.common.month.december",description:"December month translation",message:"December"})];var g=function(e){var t,a,g,h,u,b,p=e.children,v=e.frontMatter,E=e.metadata,N=e.truncated,_=e.isBlogPostPage,f=void 0!==_&&_,k=E.date,y=E.permalink,Z=E.tags,T=E.readingTime,I=v.author,P=v.title,w=v.subtitle,M=v.image,A=v.keywords,L=v.author_url||v.authorURL,D=v.author_title||v.authorTitle,J=v.author_image_url||v.authorImageURL;return n.createElement(n.Fragment,null,n.createElement(s.Z,{keywords:A,image:M}),n.createElement("article",{className:f?void 0:"margin-bottom--xl"},(t=f?"h1":"h2",a=f?"h2":"h3",g=k.substring(0,10).split("-"),h=g[0],u=d[parseInt(g[1],10)-1],b=parseInt(g[2],10),n.createElement("header",null,n.createElement(t,{className:(0,r.Z)("margin-bottom--sm",c.blogPostTitle)},f?P:n.createElement(m.Z,{to:y},P)),w&&n.createElement(a,{className:c.subtitle},w),n.createElement("div",{className:"margin-vert--md"},n.createElement("div",{className:c.heading},n.createElement("div",{className:c.headingPhoto},J&&n.createElement(m.Z,{className:"avatar__photo-link avatar__photo "+c.avatarImage,href:L},n.createElement("img",{src:J,alt:I}))),n.createElement("div",{className:c.headingDetails},n.createElement("span",null,n.createElement(m.Z,{className:c.authorName,href:L},I),", ",n.createElement("span",{className:c.authorTitle},D)),n.createElement("time",{dateTime:k,className:c.blogPostDate},n.createElement("br",null),n.createElement(l.Z,{id:"theme.blog.post.date",description:"The label to display the blog post date",values:{day:b,month:u,year:h}},"{month} {day}, {year}")," ",T&&n.createElement(n.Fragment,null," \xb7 ",n.createElement(l.Z,{id:"theme.blog.post.readingTime",description:"The label to display reading time of the blog post",values:{readingTime:Math.ceil(T)}},"{readingTime} min read")))))))),n.createElement("div",{className:"markdown"},n.createElement(i.Zo,{components:o.Z},p)),(Z.length>0||N)&&n.createElement("footer",{className:"row margin-vert--lg"},Z.length>0&&n.createElement("div",{className:c.tags},n.createElement("strong",null,n.createElement(l.Z,{id:"theme.tags.tagsListLabel",description:"The label alongside a tag list"},"Tags:")),Z.map((function(e){var t=e.label,a=e.permalink;return n.createElement(m.Z,{key:a,className:c.tag,to:a},t)}))),N&&n.createElement("div",{className:"col text--right"},n.createElement(m.Z,{to:E.permalink,"aria-label":"Read more about "+P},n.createElement("strong",null,n.createElement(l.Z,{id:"theme.blog.post.readMore",description:"The label used in blog post item excerpts to link to full blog posts"},"Read More")))))))}},8817:function(e,t,a){a.d(t,{Z:function(){return g}});var n=a(7294),r=a(6010),i=a(3699),l="sidebar_2tcx",m="sidebarItemTitle_1XYP",o="sidebarItemList_3OT0",s="sidebarItem_ic0g",c="sidebarItemLink_3PPN",d="sidebarItemLinkActive_23br";function g(e){var t=e.sidebar;return 0===t.items.length?null:n.createElement("div",{className:"blog-sidebar"},n.createElement("div",{className:(0,r.Z)(l,"thin-scrollbar")},n.createElement("h3",{className:m},t.title),n.createElement("ul",{className:o},t.items.map((function(e){return n.createElement("li",{key:e.permalink,className:s},n.createElement(i.Z,{isNavLink:!0,to:e.permalink,className:c,activeClassName:d},e.title))})))))}}}]);