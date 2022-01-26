<template>
    <div class="info">
        <div class="statistic"></div>
        <div class="brief">
            <div>
                <div class="name"></div>
                <div class="value"></div>
            </div>
            <div>
                <div class="name"></div>
                <div class="value"></div>
            </div>
            <div>
                <div class="name"></div>
                <div class="value"></div>
            </div>
            <div>
                <div class="name"></div>
                <div class="value"></div>
            </div>
        </div>
        <div class="third-party"></div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"

const name = ref<String>("Info")
</script>

<style scoped>
.info {
    width: 100%;
    height: 100%;
    grid-area: info;
    border: 1px solid silver;
    box-sizing: border-box;
    display: grid;
    grid-template-areas: "statistic brief social";
    grid-template-columns: 530px 660px 1fr;
}

.statistic {
    width: 100%;
    height: 300px;
    grid-area: statistic;
    border: 1px solid silver;
    box-sizing: border-box;
}

.brief {
    width: 100%;
    height: 300px;
    grid-area: brief;
    border: 1px solid silver;
    box-sizing: border-box;
}

.third-party {
    width: 100%;
    height: 300px;
    grid-area: social;
    border: 1px solid silver;
    box-sizing: border-box;
}

.name {
    width: 100px;
    height: 50px;
    margin-left: 20px;
    margin-top: 20px;
    border: 1px solid silver;
    box-sizing: border-box;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1.2rem;
}

.value {
    width: 510px;
    height: 50px;
    margin-top: 20px;
    margin-left: 28px;
    border: 1px solid silver;
    box-sizing: border-box;
    display: inline-block;
    vertical-align: bottom;
    font-size: 1.2rem;
}
</style>