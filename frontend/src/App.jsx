import { useEffect, useState } from "react";
import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";
import api from "./services/api";

export default function App() {
  const [auth, setAuth] = useState(null);

  useEffect(() => {
    const checkAuth = async () => {
      try {
        await api.get("/me");
        setAuth(true);
      } catch {
        setAuth(false);
      }
    };

    checkAuth();
  }, []);

  if (auth === null) return <div>Loading...</div>;

  return auth ? (
    <Dashboard onLogout={() => setAuth(false)} />
  ) : (
    <Login onLogin={() => setAuth(true)} />
  );
}
