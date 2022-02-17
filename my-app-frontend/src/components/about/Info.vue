<template>
    <div class="info">
        <PH1></PH1>
        <div class="statistic">
            <div>
                <div class="name" @click="router.push('/article_list')">你的地区</div>
                <input type="text" class="value" v-model="region" placeholder="手动输入格式(全拼): [省份/城市c]-[区/县]" @keydown.enter.down="getCoordinate"/>
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
                <div v-for="i in hourlyWeather" style="display: inline-block">
                    <img class="icon" :src="i.icon" style="display: block">
                    <div class="time">{{i.time}}时</div>
                </div>
            </div>
        </div>
        <div class="brief">
            <div>
                <div class="name article-list" @click="router.push('/article_list')">文章总数</div>
                <div class="value"></div>
            </div>
            <div>
                <div class="name">上次更新</div>
                <div class="value"></div>
            </div>
            <div>
                <div class="name">最多标签</div>
                <div class="value"></div>
            </div>
            <div>
                <div class="name">最新文章</div>
                <div class="value"></div>
            </div>
        </div>
        <PH2></PH2>
    </div>
</template>

<script setup lang="ts">
import {reactive, ref, watch} from "vue"
import {useRouter} from "vue-router";
import PH1 from "../placeholder/PH1.vue"
import PH2 from "../placeholder/PH2.vue"
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";

const name = ref<string>("Info")
const router = useRouter()

const myKey = "80dfc319835144c4a59572d5319f305b"
const iconDir = "../../icons/"

const total = ref<number>(0)
const lastedUpdate = ref<string>('')
const lastedAdd = ref<string>('')
const mostTag = ref<string>('')

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
            console.log(coordinate)
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
            console.log(data)
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
    border: 2px solid lightsalmon;
    box-sizing: border-box;
    border-radius: 20px;
}

.brief {
    width: 100%;
    height: 300px;
    grid-area: brief;
    border: 2px solid salmon;
    box-sizing: border-box;
    border-radius: 20px;
}

.name {
    width: 100px;
    height: 50px;
    margin-left: 20px;
    margin-top: 20px;
    border: 2px solid sandybrown;
    box-sizing: border-box;
    border-radius: 18px;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1.2rem;
    font-weight: bolder;
    text-align: center;
    line-height: 46px;
    background-color: rgba(255, 140, 0, 0.05);
}

.value {
    width: calc(100% - 168px);
    height: 50px;
    margin-top: 20px;
    margin-left: 20px;
    margin-right: 20px;
    padding: 0 0 0 8px;
    border: none;
    border-radius: 16px;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1.2rem;
    line-height: 50px;
    outline: none;
    background-color: rgba(255, 0, 188, 0.05);
}

.article-list {
}

.article-list:hover {
    opacity: 60%;
}

.article-list:active {
    opacity: 90%;
}

.weather {
    width: calc(100% - 160px);
    height: 50px;
    margin-top: 20px;
    margin-left: 20px;
    margin-right: 20px;
    padding: 0;
    border: none;
    border-radius: 16px;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1.2rem;
    line-height: 50px;
    outline: none;
    background-color: rgba(255, 0, 188, 0.05);
}

.weather-item {
    width: 100px;
    height: 100%;
    vertical-align: bottom;
    line-height: 50px;
    display: inline-block;
}

.future-weather {
    width: 100%;
    height: 140px;
    text-align: left;
    padding-top: 20px;
    padding-left: 18px;
}

.icon {
    width: 70px;
    height: 70px;
    margin-left: 30px;
    outline: none;
}

.time {
    width: 70px;
    height: 70px;
    margin-left: 30px;
    text-align: center;
    display: inline-block;
}
</style>