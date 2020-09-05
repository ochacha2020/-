export async function callVisonAPI(image) {
    // Base64でエンコードした画像を渡す
    const imageArray = image.split(",");
    if (imageArray.length != 2) {
      console.log("error");
      return;
    }

    const base_img = imageArray[1]; // 最初のbase64; image...みたいなやつを消す

    const res = await fetch("http://127.0.0.1:8080/api/v1/vision", {
      method: "POST",
      body: JSON.stringify({
        base_img: base_img,
      }),
    });
    const json = await res.json();
    return json
  }

