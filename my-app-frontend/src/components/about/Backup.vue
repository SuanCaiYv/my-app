<template>
    <div class="backup">
        <button class="button" @click="backup">备份数据</button>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {BASE_URL, httpClient} from "../../net";
import {Response} from "../../common/interface";
import storage from "../../util/storage";
import {Constant} from "../../common/systemconstant";

const name = ref<string>("Backup")

const backup = function () {
    // 创建a标签
    const link = document.createElement('a')
    // 设置a标签的href（点击地址）
    link.href = BASE_URL + "/backup/" + storage.get(Constant.ACCESS_TOKEN)
    // 设置a标签属性
    link.setAttribute('download', new Date().toLocaleString() + ".json")
    // 点击a标签
    document.body.appendChild(link)
    link.click()
    // 移除a标签
    document.body.removeChild(link)
}
</script>

<style scoped>
.backup {
    width: 100%;
    height: 100%;
    grid-area: backup;
    text-align: left;
}

.button {
    width: auto;
    height: 40px;
    margin-left: 0;
    margin-top: 10px;
    margin-bottom: 10px;
    padding: 0 20px;
    border: none;
    box-sizing: border-box;
    border-radius: 18px;
    font-size: 1.0rem;
    font-weight: bolder;
    background-color: rgba(0,0,0,0.1);
}

.button:hover {
    background-color: rgba(0,0,0,0.15);
}

.button:active {
    background-color: rgba(0,0,0,0.2);
}
</style>