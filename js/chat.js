window.onload = function () {
    const reqBox = document.querySelector('#question-box');
    const vh = Math.max(document.documentElement.clientHeight || 0, window.innerHeight || 0);
    function autoResize() {
        if (reqBox.scrollHeight > vh * 0.3) {
            reqBox.style.height = 'auto';
            reqBox.style.height = reqBox.scrollHeight + 'px';
        } else {
            reqBox.style.height = '30vh';
        }
    }
    reqBox.addEventListener('input', autoResize);

    const button = document.querySelector('#submit-button');
    const loadBox = document.getElementById("loading-box");
    button.addEventListener('click', () => {
        loadBox.getElementsByClassName("loading-word")[0].textContent = "获取数据中";
        loadBox.className = "";
        loadBox.style.display = 'block';

        let reqText = reqBox.value;
        reqBox.value = "";

        let xmlHttp = new XMLHttpRequest();
        xmlHttp.open("POST", "/GetData", true);
        xmlHttp.send(JSON.stringify({ Text: reqText }));
        xmlHttp.onreadystatechange = function () {
            if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
                loadBox.getElementsByClassName("loading-word")[0].textContent = "解析数据中";
                ShowAnswer(JSON.parse("[{\"content\":\"" + reqText + "\"}," + xmlHttp.responseText + "]"));
            } else {
                loadBox.getElementsByClassName("loading-word")[0].textContent = "获取数据失败";
            }
        }
    })

    //获取历史记录
    loadBox.getElementsByClassName("loading-word")[0].textContent = "获取数据中";
    loadBox.className = "";
    loadBox.style.display = 'block';
    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open("POST", "/GetHistory", true);
    xmlHttp.send(null);
    xmlHttp.onreadystatechange = function () {
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
            loadBox.getElementsByClassName("loading-word")[0].textContent = "解析数据中";
            const ansBox = document.getElementById("answer-box");
            ansBox.innerText = xmlHttp.responseText;
            setTimeout(function () {
                loadBox.className = "loaded";
                loadBox.style.display = 'none';
            }, 0);
        } else {
            loadBox.getElementsByClassName("loading-word")[0].textContent = "获取数据失败";
        }
    }
}

function ShowAnswer(data) {
    console.log(data);
    const ansBox = document.getElementById("answer-box");
    if (ansBox.innerText != "")
        ansBox.innerText += "\n";
    ansBox.innerText += "你：" + data[0].content + "\n" + "ChatGpt：" + data[1].content;

    const loadBox = document.getElementById("loading-box");
    setTimeout(function () {
        loadBox.className = "loaded";
        loadBox.style.display = 'none';
    }, 0);
}