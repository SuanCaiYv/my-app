<template>
    <div class="kind-or-tag">
        <div class="value">{{ props.name }}</div>
        <div class="delete" @click="del">✖️</div>
    </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
import {httpClient} from "../../net";
import {Response} from "../../common/interface";
import alertFunc from "../../util/alert";
import {confirmFunc} from "../../util/confirm";

const name = ref<string>("KindOrTagItem")

const props = defineProps({
    id: String,
    name: String,
    kind: String,
    onDeleted: Function,
})

const del = function () {
    confirmFunc("确认删除?", function () {}, function () {
        httpClient.delete("/article/" + props.kind + "/" + props.id, {}, true, function (resp: Response) {
            if (!resp.ok) {
                console.log(resp.errMsg)
            } else {
                alertFunc("删除成功", function () {})
                // @ts-ignore
                props.onDeleted(props.id)
            }
        })
    })
}
</script>

<style scoped>
.kind-or-tag {
    width: auto;
    height: 30px;
    margin: 10px 5px;
    border-radius: 16px;
    display: inline-block;
    background-color: rgba(0,0,0,0.25);
}

.value {
    width: auto;
    height: 100%;
    border-radius: 16px;
    padding-left: 8px;
    padding-right: 4px;
    display: inline-block;
    text-align: left;
    line-height: 30px;
    font-size: 1.2rem;
    font-weight: bolder;
    vertical-align: bottom;
}

.delete {
    width: 30px;
    height: 100%;
    border: none;
    border-radius: 16px;
    display: inline-block;
    font-size: 1.2rem;
    text-align: center;
    line-height: 30px;
    vertical-align: bottom;
}

.delete:hover {
    background-color: rgba(0,0,0,0.15);
}

.delete:active {
    background-color: rgba(0,0,0,0.3);
}
</style>