<template>
  <div class="echarts">
    <ECharts :option="option" />
  </div>
</template>

<script setup lang="ts" name="pie">
import { ECOption } from "@/components/ECharts/config";
import ECharts from "@/components/ECharts/index.vue";
import { ref, onMounted } from "vue";
import axios from "axios";

// 使用 ref 创建一个响应式引用
const pieData = ref([
  { value: 0, name: "Running" },
  { value: 0, name: "Driving" },
  { value: 0, name: "Cycling" }
]);

// 定义获取数据的函数
const fetchData = async () => {
  try {
    const response = await axios.get("http://120.46.81.37:8080/admin/index");
    // 返回的数据格式为 [{ value: 20, name: "Running" }, ...]
    pieData.value = response.data;
  } catch (error) {
    console.error("Error fetching data:", error);
  }
};

// 当组件挂载后，获取数据
onMounted(fetchData);

const option: ECOption = {
  title: {
    text: "Track mode",
    subtext: "percentage",
    left: "56%",
    top: "45%",
    textAlign: "center",
    textStyle: {
      fontSize: 18,
      color: "#767676"
    },
    subtextStyle: {
      fontSize: 15,
      color: "#a1a1a1"
    }
  },
  tooltip: {
    trigger: "item"
  },
  legend: {
    top: "4%",
    left: "2%",
    orient: "vertical",
    icon: "circle", //图例形状
    align: "left",
    itemGap: 20,
    textStyle: {
      fontSize: 13,
      color: "#a1a1a1",
      fontWeight: 500
    },
    formatter: function (name: string) {
      let dataCopy = "";
      for (let i = 0; i < pieData.value.length; i++) {
        if (pieData.value[i].name == name && pieData.value[i].value >= 10000) {
          dataCopy = (pieData.value[i].value / 10000).toFixed(2);
          return name + "      " + dataCopy + "w";
        } else if (pieData.value[i].name == name) {
          dataCopy = pieData.value[i].value + "";
          return name + "      " + dataCopy;
        }
      }
      return "";
    }
  },
  series: [
    {
      type: "pie",
      radius: ["70%", "40%"],
      center: ["57%", "52%"],
      silent: true,
      clockwise: true,
      startAngle: 150,
      data: pieData.value,
      labelLine: {
        length: 80,
        length2: 30,
        lineStyle: {
          width: 1
        }
      },
      label: {
        position: "outside",
        show: true,
        formatter: "{d}%",
        fontWeight: 400,
        fontSize: 19,
        color: "#a1a1a1"
      },
      color: [
        {
          type: "linear",
          x: 0,
          y: 0,
          x2: 0.5,
          y2: 1,
          colorStops: [
            {
              offset: 0,
              color: "#feb791" // 0% 处的颜色
            },
            {
              offset: 1,
              color: "#fe8b4c" // 100% 处的颜色
            }
          ]
        },
        {
          type: "linear",
          x: 0,
          y: 0,
          x2: 0.5,
          y2: 1,
          colorStops: [
            {
              offset: 0,
              color: "#77DDFF" // 0% 处的颜色
            },
            {
              offset: 1,
              color: "#0066FF" // 100% 处的颜色
            }
          ]
        },
        {
          type: "linear",
          x: 0,
          y: 0,
          x2: 1,
          y2: 0.5,
          colorStops: [
            {
              offset: 0,
              color: "#b898fd" // 0% 处的颜色
            },
            {
              offset: 1,
              color: "#8347fd" // 100% 处的颜色
            }
          ]
        }
      ]
    }
  ]
};
</script>

<style lang="scss" scoped>
.echarts {
  width: 100%;
  height: 100%;
}
</style>
