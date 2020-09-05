export function getEmocon(json_res){
  var feeling="";
  if (json_res["surprise"] > 50) feeling = "surprised";
  else if (json_res["sorrow"] > 50) feeling = "sad";
  else if (json_res["anger"] > 50) feeling = "angry";
  else if (json_res["joy"] > 70) feeling = "perfectly_joy";
  else if (json_res["joy"] > 50) feeling = "soso_joy";
  return feeling;
}
