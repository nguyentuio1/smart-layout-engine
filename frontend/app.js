// 1. Cấu hình WebAssembly
const go = new Go();
const importObject = go.importObject;
// Sửa lỗi "gojs" để tương thích với Go 1.25+
importObject.gojs = go.importObject.go; 

let items = []; // Mảng lưu trữ danh sách đồ đạc người dùng nhập

// Tải file Wasm
WebAssembly.instantiateStreaming(fetch("main.wasm"), importObject).then((result) => {
    go.run(result.instance);
    document.getElementById("status").innerText = "Sẵn sàng!";
}).catch((err) => {
    console.error("Lỗi nạp Wasm:", err);
    document.getElementById("status").innerText = "Lỗi nạp Wasm";
});

// 2. Logic thêm vật thể vào danh sách
document.getElementById("addBtn").onclick = () => {
    const nameInput = document.getElementById("itemName");
    const wInput = document.getElementById("itemW");
    const hInput = document.getElementById("itemH");

    const name = nameInput.value.trim();
    const w = parseInt(wInput.value);
    const h = parseInt(hInput.value);

    if (!name || isNaN(w) || isNaN(h)) {
        alert("Vui lòng nhập đầy đủ Tên, Rộng và Dài!");
        return;
    }

    // Lưu vào mảng dữ liệu
    items.push({ id: name, size: { w: w, h: h } });

    // Hiển thị lên giao diện (UL)
    const ul = document.getElementById("itemList");
    const li = document.createElement("li");
    li.innerHTML = `<strong>${name}</strong>: ${w}px x ${h}px`;
    ul.appendChild(li);

    // Reset ô nhập
    nameInput.value = "";
    wInput.value = "";
    hInput.value = "";
    nameInput.focus();
};

// 3. Logic gọi bộ giải Go và Vẽ kết quả
document.getElementById("solveBtn").onclick = () => {
    if (items.length === 0) {
        alert("Danh sách đang trống! Hãy thêm đồ đạc trước.");
        return;
    }

    try {
        // Chuyển mảng sang JSON và gửi cho Go
        const inputJSON = JSON.stringify(items);
        const resultJSON = solveInGo(inputJSON);
        const result = JSON.parse(resultJSON);

        if (result.error) {
            alert("Thuật toán báo lỗi: " + result.error);
            return;
        }

        renderLayout(result);
    } catch (e) {
        console.error("Lỗi thực thi:", e);
        alert("Có lỗi xảy ra khi tính toán. Kiểm tra Console!");
    }
};

// 4. Hàm vẽ lên Canvas
function renderLayout(placedItems) {
    const canvas = document.getElementById("roomCanvas");
    const ctx = canvas.getContext("2d");

    // Xóa trắng Canvas
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    placedItems.forEach((item, index) => {
        // Tạo màu sắc ngẫu nhiên dễ nhìn
        const hue = (index * 137.5) % 360; 
        ctx.fillStyle = `hsla(${hue}, 70%, 80%, 0.8)`;
        ctx.strokeStyle = `hsla(${hue}, 70%, 30%, 1)`;
        ctx.lineWidth = 2;

        const { x, y } = item.position;
        const { w, h } = item.size;

        // Vẽ khối vật thể
        ctx.fillRect(x, y, w, h);
        ctx.strokeRect(x, y, w, h);

        // Vẽ tên vật thể
        ctx.fillStyle = "#333";
        ctx.font = "bold 12px sans-serif";
        ctx.fillText(item.id, x + 5, y + 18);
    });
}
