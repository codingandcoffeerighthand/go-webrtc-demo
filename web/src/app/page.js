"use client";
import { useEffect, useState } from "react";
import Web3 from "web3";
export default function Home() {
  const [account, setAccount] = useState("");
  const [status, setStatus] = useState("");
  const [isConnected, setIsConnected] = useState(false);
  
  const ping = async () => {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/ping`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });
      const data = await response.json();
      console.log(data);
      setIsConnected(true);
    }
    catch (error) {
      console.error("Error pinging server:", error);
      setIsConnected(false);
    }
  }
  
  useEffect(() => {
    ping();
  }, []);
  
  const login = async () => {
    try {
      const web3 = new Web3(window.ethereum);
      await window.ethereum.request({ method: "eth_requestAccounts" });
      const accounts = await web3.eth.getAccounts();
      const selectedAccount = accounts[0];
      setAccount(selectedAccount);
      const nonce = `#nonce ${selectedAccount}`;
      const signature = await web3.eth.sign(
        nonce,
        selectedAccount,
        ""
      ); 
      console.log("Signature:", signature);
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ 
          wallet_address: selectedAccount,
          signature: signature,
          nonce: nonce, 
         }),
      });
      const data = await response.json();
      setStatus(data.token);
    } catch (error) {
      console.error("Error logging in:", error);
    }
  }

  return (
    <div>
    <p>Status: {`${isConnected}`} </p>
      <button onClick={login}>Login</button>
      <p> Account: {account}</p>
      <p> Status: {status}</p>
    </div> 
  );
}
