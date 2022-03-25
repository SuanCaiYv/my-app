<template>
    <div class="info">
        <PH1></PH1>
        <div class="statistic">
            <div>
                <div class="name" @click="router.push('/article_list')">你的地区</div>
                <input type="text" class="value indicate" v-model="region" placeholder="手动输入格式(全拼): [省份/城市c]-[区/县]" @keydown.enter.down="getCoordinate"/>
            </div>
            <div>
                <div class="name">实况天气</div>
                <div class="weather">
                    <div class="weather-item">温度: {{temp}}</div>
                    <div class="weather-item">体感: {{bodyFeel}}</div>
                    <div class="weather-item">{{desc}}</div>
                    <img class="weather-item" :src="iconPath"/>
                </div>
            </div>
            <div class="future-weather">
                <div style="width: 10px; height: 50px; display: inline-block"></div>
                <div v-for="i in hourlyWeather" style="display: inline-block">
                    <img class="icon" :src="i.icon" style="display: block">
                    <div class="time">{{i.time}}时</div>
                </div>
            </div>
        </div>
        <div class="brief">
            <div>
                <div class="name article-list" @click="router.push('/article_list')">文章总数</div>
                <div class="value">{{total}}</div>
            </div>
            <div>
                <div class="name">上次更新</div>
                <div class="value">{{lastedUpdate}}</div>
            </div>
            <div>
                <div class="name">最新标签</div>
                <div class="value">{{ newestTag }}</div>
            </div>
            <div>
                <div class="name">最新文章</div>
                <div class="value">{{lastedAdd}}</div>
            </div>
        </div>
        <PH2></PH2>
    </div>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue"
import {useRouter} from "vue-router";
import PH1 from "../placeholder/PH1.vue"
import PH2 from "../placeholder/PH2.vue"
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";
import {httpClient} from "../../net";
import {Response} from "../../common/interface";
import alertFunc from "../../util/alert";
import {parseBoolean, toListResult} from "../../util/base";

const name = ref<string>("Info")
const router = useRouter()

const myKey = "80dfc319835144c4a59572d5319f305b"
const iconDir = "../../icons/"

const total = ref<number>(0)
const lastedUpdate = ref<string>('暂无更新')
const lastedAdd = ref<string>('暂无文章')
const newestTag = ref<string>('暂无标签')

const region = ref<string>('')
let coordinate = ''
const temp = ref<string>('')
const bodyFeel = ref<string>('')
const iconPath = ref<string>('')
const desc = ref<string>('')
const hourlyWeather = reactive<Array<any>>([])

const getCity = function () {
    fetch("https://ipapi.co/json")
        .then(function (response) {
            return response.json();
        })
        .then(function (data) {
            region.value = data.city + "-" + data.region
            getCoordinate()
        })
}

const getCoordinate = function () {
    storage.set(Constant.LOCAL_CITY, region.value)
    let ss = region.value.split("-")
    let url = "https://geoapi.qweather.com/v2/city/lookup?" + "key=" + myKey + "&location=" + ss[1] + "&adm=" + ss[0]
    fetch(url)
        .then(function (resp) {
            return resp.json()
        })
        .then(function (data) {
            coordinate = data.location[0].lon + "," + data.location[0].lat
            storage.set(Constant.LOCAL_COORDINATE, coordinate)
            getWeather()
        })
}

const getWeather = function () {
    let url = "https://devapi.qweather.com/v7/weather/now?" + "key=" + myKey + "&location=" + storage.get(Constant.LOCAL_COORDINATE)
    fetch(url)
        .then(function (resp) {
            return resp.json()
        })
        .then(function (data) {
            if (data.code !== "200") {
                region.value = "地址出错啦！"
                return
            }
            temp.value = data.now.temp
            bodyFeel.value = data.now.feelsLike
            iconPath.value = iconDir + data.now.icon + ".svg"
            desc.value = data.now.text
            getHourlyWeather()
        })
}

const getHourlyWeather = function () {
    let url = "https://devapi.qweather.com/v7/weather/24h?" + "key=" + myKey + "&location=" + storage.get(Constant.LOCAL_COORDINATE)
    fetch(url)
        .then(function (resp) {
            return resp.json()
        })
        .then(function (data) {
            hourlyWeather.splice(0, hourlyWeather.length)
            for (let i = 0; i < 5; ++ i) {
                hourlyWeather.push({
                    icon: iconDir + data.hourly[i].icon + ".svg",
                    time: new Date(data.hourly[i].fxTime).getHours()
                })
            }
        })
}

if (storage.get(Constant.LOCAL_COORDINATE) === "") {
    getCity()
} else {
    region.value = storage.get(Constant.LOCAL_CITY)
    getWeather()
}

const fetchLatestArticle = function () {
    httpClient.get("/article/list", {
        sort: "release_time",
        desc: false,
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        } else {
            const list = toListResult(resp.data)
            total.value = list.total
            if (list.list.length > 0) {
                lastedAdd.value = "《" + list.list[0].article_name + "》"
            }
        }
    })
}
const fetchLatestUpdate = function () {
    httpClient.get("/article/list", {
        sort: "updated_time",
        desc: false,
    }, true, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        } else {
            const list = toListResult(resp.data)
            if (list.list.length > 0) {
                lastedUpdate.value = "《" + list.list[0].article_name + "》"
            }
        }
    })
}
const fetchNewestTag = function () {
    httpClient.get("/article/tag_list", {
        page_num: -1,
    }, false, function (resp: Response) {
        if (!resp.ok) {
            alertFunc(resp.errMsg, function () {})
        } else {
            const list = toListResult(resp.data)
            if (list.list.length > 0) {
                newestTag.value = list.list[0].tag_name
            }
        }
    })
}

if (parseBoolean(storage.get(Constant.AUTHENTICATED))) {
    fetchLatestArticle()
    fetchLatestUpdate()
    fetchNewestTag()
}
</script>

<style scoped>
.info {
    width: 100%;
    height: 100%;
    grid-area: info;
    display: grid;
    grid-template-areas: "ph1 statistic brief ph2";
    grid-template-columns: 1fr 570px 570px 1fr;
}

.statistic {
    width: 100%;
    height: 300px;
    grid-area: statistic;
    border: none;
    box-sizing: border-box;
    border-radius: 20px;
    text-align: left;
}

.brief {
    width: 100%;
    height: 300px;
    grid-area: brief;
    border: none;
    box-sizing: border-box;
    border-radius: 20px;
    text-align: left;
}

.name {
    width: 100px;
    height: 50px;
    margin-left: 20px;
    margin-top: 20px;
    border: none;
    box-sizing: border-box;
    border-radius: 18px;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1.2rem;
    font-weight: bolder;
    text-align: center;
    line-height: 50px;
    background-color: white;
}

.value {
    max-width: calc(100% - 176px);
    min-width: 30px;
    height: 50px;
    margin-top: 20px;
    margin-left: 20px;
    margin-right: 20px;
    padding: 0 8px 0 8px;
    border: none;
    border-radius: 16px;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1.2rem;
    font-weight: bolder;
    line-height: 50px;
    outline: none;
    overflow-x: auto;
    overflow-y: hidden;
    background-color: rgba(0,0,0,0.05);
}

.article-list {
    background-color: rgba(0,0,0,0.05);
}

.article-list:hover {
    background: rgba(0,0,0,0.15);
}

.article-list:active {
    background-color: rgba(0,0,0,0.25);
}

.weather {
    width: calc(100% - 176px);
    height: 50px;
    margin-top: 20px;
    margin-left: 20px;
    margin-right: 20px;
    padding: 0 8px 0 8px;
    border: none;
    border-radius: 16px;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1.2rem;
    line-height: 50px;
    outline: none;
    background-color: rgba(0,0,0,0.05);
}

.weather-item {
    width: 96px;
    height: 100%;
    vertical-align: bottom;
    line-height: 50px;
    font-weight: bolder;
    display: inline-block;
}

.future-weather {
    width: 100%;
    height: 140px;
    text-align: left;
    padding-top: 30px;
    padding-left: 18px;
}

.icon {
    width: 70px;
    height: 70px;
    margin-right: 42px;
    outline: none;
}

.time {
    width: 70px;
    height: 70px;
    margin-right: 42px;
    text-align: center;
    font-weight: bolder;
    display: inline-block;
}

.indicate {
}

.indicate:hover {
    background-color: lightgray;
}
</style>