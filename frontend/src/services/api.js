import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080",
});

// Ambil token dari localStorage saat app pertama kali load
const token = localStorage.getItem("token");
if (token) {
  api.defaults.headers.common["Authorization"] = `Bearer ${token}`;
}

export const setToken = (token) => {
  localStorage.setItem("token", token);
  api.defaults.headers.common["Authorization"] = `Bearer ${token}`;
};

export const clearToken = () => {
  localStorage.removeItem("token");
  delete api.defaults.headers.common["Authorization"];
};

export default api;
