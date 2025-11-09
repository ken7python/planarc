<script setup lang="ts">
  import { ref, onMounted } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
  import { todoModule } from "../logic/todo";
  import { mic } from '@/logic/mic';
  import { subjectModule } from "@/logic/subject";
  import { statusModule } from "@/logic/status";
  import ToDoCard from "../components/ToDoCard.vue";

  import { CONST } from '@/logic/const';

  import Addicon from '@/assets/icons/add.svg';
  import MicIcon from '@/assets/icons/mic.svg';

  let subjectName = ref('');
  let todoText = ref('');
  let status = ref('MUST');

  import mode1 from '@/assets/icons/1.svg';
  import mode2 from '@/assets/icons/2.svg';
  import mode3 from '@/assets/icons/3.svg';
  import mode4 from '@/assets/icons/4.svg';

  import saveicon from '@/assets/icons/save.svg';

  const props = defineProps({
    date: String
  })

  const communication_loading = ref<boolean>(false);
  const communication_saving = ref<boolean>(false);

  const TODO_MUST = ref<any[]>([]);
  const TODO_WANT = ref<any[]>([]);
  const TODO_checked = ref<any[]>([]);

  let subjects = ref<any[]>([]);

  async function loadData() {
    const subject_list = await subjectModule.getList();
    //console.log(subject_list);
    subjects.value = subject_list;

    const ToDoList = await todoModule.getListGroup(props.date);
    //console.log(await ToDoList);

    TODO_MUST.value = await ToDoList.MUST;
    TODO_WANT.value = await ToDoList.WANT;
    TODO_checked.value = await ToDoList.checked;

    //console.log("MUST---");
    //console.log(TODO_MUST.value);
    communication_loading.value = false;
  }

  async function add() {
    communication_saving.value = true;
    await todoModule.add(props.date, todoText.value, subjectName.value, status.value);
    todoText.value = "";
    communication_saving.value = false;
    loadData();
  };


  const tempStr = ref<string>('');
  const tempmode =ref<string>('');

  const micbtn = async () => {
    if (!mic.shouldRestart.value) {        // ← .value
      console.log("start");
      mic.result = '';
      // mic.micON.value = true;
      await mic.start();
    } else {
      console.log("stop");
      // mic.micON.value = false;
      mic.stop();
    }
  };

  // onresult は setup後に一度だけ付ける
  mic.rec && (mic.rec.onresult = (e) => {
    let interim = '';
    for (let i = e.resultIndex; i < e.results.length; i++) {
      const transcript = e.results[i][0].transcript;
      if (e.results[i].isFinal) {
        mic.result = transcript;
        mic.tmp = '';
      } else {
        interim += transcript;
        mic.tmp = interim;
      }
    }

    tempStr.value = mic.result || mic.tmp;
    if (tempmode.value === 'todo') {
      todoText.value = tempStr.value;
    } else if (tempmode.value === 'enjoyment') {
      enjoyment.value = tempStr.value;
    }
  });

  function micTodo() {
    tempmode.value = 'todo';
    micbtn();
  }
  function micEnjoyment() {
    tempmode.value = 'enjoyment';
    micbtn();
  }

  let enjoyment = ref<string>('');

  async function setEnjoyment(enjoy: string) {
    //console.log(enjoy);
    await statusModule.setEnjoyment(props.date,enjoy);
    loadStatus();
  }

  let selectedMood = ref<number>(null);
  async function moodSelect(mood: number) {
    // selectedMood.value = mood;
    await statusModule.setMood(props.date,mood);
    loadStatus();
  }
  const getMoodStyle = (mood: number) => {
    return {
      color: selectedMood.value === mood ? '#4d5161' : 'gray',
      background: selectedMood.value === mood ? 'yellow' : 'white',
      cursor: 'pointer'
    };
  };


  async function loadStatus() {
    const status = await statusModule.get(props.date);
    //console.log(status);
    if (status) {
      enjoyment.value = status.Enjoyment || '';
      selectedMood.value = status.Mood || null;
    } else {
      enjoyment.value = '';
      selectedMood.value = null;
    }
  }

  communication_loading.value = true;
  loadData();
  loadStatus();

  // test

  const publicVapidKey = "BKdgyFaYbmA8NNQvlHbr6TQ6wJudtWWzmlcDmPogbp9ppkRuvB7kQThDjVw0LDwjynesVAQvlRlFkdfMu45KO6g";

  async function register() {
    console.log("=== プッシュ通知登録開始 ===");

    try {
      if ("serviceWorker" in navigator) {
        console.log("1. Service Worker対応確認 OK");

        // ビルド後に正しいパスで登録
        const swPath = import.meta.env.PROD ? '/sw.js' : `${import.meta.env.BASE_URL}sw.js`;
        console.log("Service Worker登録パス:", swPath);

        const reg = await navigator.serviceWorker.register(swPath);
        console.log("2. Service Worker 登録完了:", reg);

        // 通知許可の確認
        console.log("3. 通知許可確認開始");
        const permission = await Notification.requestPermission();
        console.log("4. 通知許可結果:", permission);

        if (permission !== "granted") {
          console.error("❌ 通知が許可されていません");
          alert("通知が許可されていません。ブラウザの設定で通知を許可してください。");
          return;
        }

        // プッシュマネージャーの確認
        if (!reg.pushManager) {
          console.error("❌ プッシュマネージャーが利用できません");
          alert("このブラウザはプッシュ通知に対応していません");
          return;
        }

        console.log("5. プッシュ購読開始");
        console.log("VAPIDキー:", publicVapidKey.substring(0, 20) + "...");

        const sub = await reg.pushManager.subscribe({
          userVisibleOnly: true,
          applicationServerKey: urlBase64ToUint8Array(publicVapidKey),
        });
        console.log("6. プッシュ購読完了:", sub);
        console.log("購読エンドポイント:", sub.endpoint);

        console.log("7. サーバー送信開始");
        const apiUrl = `${CONST.api()}/send`;
        console.log("API URL:", apiUrl);

        const response = await fetch(apiUrl, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(sub),
        });

        console.log("8. サーバーレスポンス:", response.status, response.statusText);

        if (response.ok) {
          console.log("✅ 通知登録完了");
          alert("通知登録完了🎉");
        } else {
          const errorText = await response.text();
          console.error("❌ サーバーエラー:", response.status, errorText);
          alert(`サーバーエラー: ${response.status}\n詳細: ${errorText}`);
        }
      } else {
        console.error("❌ Service Worker非対応");
        alert("このブラウザはService Workerに対応していません");
      }
    } catch (error) {
      console.error("❌ 通知登録エラー:", error);
      console.error("エラー詳細:", error.stack);
      alert(`エラーが発生しました: ${error.message}`);
    }
  }

  function urlBase64ToUint8Array(base64String) {
    const padding = "=".repeat((4 - (base64String.length % 4)) % 4);
    const base64 = (base64String + padding).replace(/-/g, "+").replace(/_/g, "/");
    const rawData = atob(base64);
    return Uint8Array.from([...rawData].map((c) => c.charCodeAt(0)));
  }

  // デバッグ用テスト通知機能
  async function testNotification() {
    console.log("=== テスト通知開始 ===");

    try {
      // 1. 直接通知テスト（Service Worker経由なし）
      if (Notification.permission === 'granted') {
        console.log("1. 直接通知テスト");
        new Notification("直接通知テスト", {
          body: "これは直接通知です（Service Worker経由なし）",
          icon: "/pwa-192x192.png"
        });

        // 2秒後にService Worker経由のテストも実行
        setTimeout(async () => {
          console.log("2. Service Worker経由テスト");

          // Service Worker経由の通知
          const reg = await navigator.serviceWorker.getRegistration();
          if (reg && reg.active) {
            await reg.showNotification("Service Worker テスト", {
              body: "これはService Worker経由の通知です",
              icon: "/pwa-192x192.png",
              tag: "test"
            });
            console.log("Service Worker通知送信完了");
          } else {
            console.error("Service Workerが見つかりません");
          }
        }, 2000);

        alert("テスト通知を送信しました！");
      } else {
        alert("通知許可が必要です");
      }
    } catch (error) {
      console.error("テスト通知エラー:", error);
      alert("テスト通知エラー: " + error.message);
    }
  }
</script>

<template>
  <div id="ToDoPage">
    <div id="top-menu">
    <div id="statusMenu">
        <details class="frame">
          <summary>今日の楽しみ</summary>
          <div>
            <div>
              <!--        <h3 style="line-height: 0;margin-bottom: 0">今日の楽しみ</h3>-->
              <div style="display: flex; align-items: center; justify-content: center;">
                <div class="micdiv">
                  <input type="text" placeholder="今日の楽しみを入力" style="width: 70vw" v-model="enjoyment" />
                  <mic-icon :style="mic.micStyle(tempmode === 'enjoyment')" @click="micEnjoyment"></mic-icon>
                </div>
                <button class="squareBtn btnSave" style="margin-left: 20px;" @click="setEnjoyment(enjoyment)"><saveicon></saveicon></button>
              </div>
            </div>
          </div>
        </details>
        <details class="frame">
          <summary>今日の気分</summary>
          <div>
            <div id="feeling">
              <!--          <h3 style="line-height: 0;margin-bottom: 0">今日の気分</h3>-->
              <div id="feeling-icons">
                <mode4 @click="moodSelect(4)" :style="getMoodStyle(4)"></mode4>
                <mode3 @click="moodSelect(3)" :style="getMoodStyle(3)"></mode3>
                <mode2 @click="moodSelect(2)" :style="getMoodStyle(2)"></mode2>
                <mode1 @click="moodSelect(1)" :style="getMoodStyle(1)"></mode1>
              </div>
            </div>
          </div>
        </details>
      </div>
    <div>
      <details class="frame">
        <summary>リスト作成</summary>
        <div id="AddTodo">
          <select class="selectbox" :style="selectStyle.getSelectStyle(subjectName)" v-model="subjectName">
            <option disabled value="">科目を選択</option>
            <option v-for="subject in subjects" :key="subject.value" :value="subject.ID">
              {{ subject.Name }}
            </option>
          </select>
          <div class="micdiv">
            <input type="text" placeholder="ToDoを入力" v-model="todoText" />
            <MicIcon class="mic" @click="micTodo" :style="mic.micStyle(tempmode === 'todo')"></MicIcon>
          </div>
          <select class="selectbox" :style="selectStyle.getSelectStyle(status)" v-model="status">
            <option value="MUST">MUST</option>
            <option value="WANT">WANT</option>
          </select>
          <br>
          <br>
          <button class="btn" style="margin: 0 auto;" @click="add" :disabled="communication_loading">
            <Addicon v-if="!communication_saving"></Addicon>
            <span v-if="!communication_saving">追加</span>

            <span v-if="communication_saving" >通信中</span>
          </button>
        </div>
        <br>
      </details>
    </div>

    </div>
    <div id="List" v-if="!communication_loading">
      <div>
        <div id="MUST" class="frame">
          <span>MUST</span>
          <span v-if="TODO_MUST.length > 0">({{ TODO_MUST.length }})</span>
          <ToDoCard :LIST="TODO_MUST" @reload="loadData" :key="TODO_MUST"/>
        </div>


        <div id="WANT" class="frame">
          <span>WANT</span>
          <span v-if="TODO_WANT.length > 0">({{ TODO_WANT.length }})</span>
          <ToDoCard :LIST="TODO_WANT" @reload="loadData" :key="TODO_WANT" />
        </div>


        <div id="Checked" class="frame">
          <span>完了</span>
          <span v-if="TODO_checked.length > 0">({{ TODO_checked.length }})</span>
          <ToDoCard :LIST="TODO_checked" @reload="loadData" :key="TODO_checked" />
        </div>
      </div>
    </div>
    <div v-else style="color: white;text-align: center; margin-top: 20px;">
      <p>通信中...</p>
    </div>

    <button @click="register" style="position: fixed; bottom: 80px; right: 20px; z-index: 1000;">
      🔔 通知を許可して送信
    </button>

    <!-- デバッグ用テスト通知ボタン -->
    <button @click="testNotification" style="position: fixed; bottom: 140px; right: 20px; z-index: 1000; background: orange; color: white; padding: 8px; border: none; border-radius: 4px;">
      🧪 テスト通知
    </button>
  </div>
</template>

<style scoped>
  #AddTodo {
    text-align: center;
  }
  #ToDoPage {
    /* background-color: #3d7fe0; */
    height: calc(100dvh - 80px - 34px - 40px - 10px);
    display: grid;
    grid-template-rows: auto 1fr;
  }

  #top-menu {
    margin-top: 10px;
    margin-bottom: 10px;
    margin-right: 0px;
    margin-left: 0px;
  }

  .selectbox,input[type="text"] {
    padding-bottom: 4px;
    margin-bottom: 0px;
    margin-top: 0px;
  }

  .frame {
    border: 2px solid var(--border-color);
    background: var(--white-color);
    border-radius: 8px;
    margin: 5px;
    padding: 5px;
  }

  #feeling-icons {
    display: flex;
    justify-content: space-around;
    align-items: center;
    height: 50px;
  }

  h3 {
    line-height: 0.75;
  }
</style>
