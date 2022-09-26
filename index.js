const site = require('./config/siteInfo');
const fs = require('fs');
const path = require('path');
const https = require('https');
const http = require('http');
const convert = require('webp-batch-convert');


if(!fs.existsSync(path.resolve(__dirname, './profile'))){
    fs.mkdirSync(path.resolve(__dirname, './profile'));
    console.log("成功创建了profile文件夹");
}

if(!fs.existsSync(path.resolve(__dirname, './webp'))){
    fs.mkdirSync(path.resolve(__dirname, './webp'));
    console.log("成功创建了webp文件夹");
}

var links = [];

const getProfile = async (biliId, key) => {
    https.get(`https://xysbtn.xiaoblogs.cn/userinfo?mid=${biliId}`, (res) => {
        res.on("data", (data) => {
            links.push({uid: biliId, link: JSON.parse(data.toString()).data.face});
        })
        res.on("end", () => {
            key += 1
            if(key === site.site.supports.length){
                console.log("获取了所有的profile data, links的长度为：", links.length);
                getJPG();
                return;
            }
            getProfile(site.site.supports[key].uid, key)
        })
    }).on('error', ()=>{})
}

const getJPG = () => {
    links.map((v, k) => {
        http.get(v.link, (res) => {
            let imgdata = "";
            res.setEncoding("binary");
            res.on("data", (data) => {
                imgdata += data;
            })
            res.on("end", () => {
                fs.writeFileSync(path.resolve(__dirname, './profile', `${v.uid}.jpg`), imgdata, "binary", (err) => {
                    if (err){
                        throw err;
                    }
                })
            })
        }).on("error", () => {})
    })
}

const mainFunction = () => {
    var key = 0
    getProfile(site.site.supports[key].uid, key)
    setTimeout(() => {
        convert.cwebp('./profile', './webp', {
            q: 80,
            quiet: true
        })
    }, 30000)
}

mainFunction();
