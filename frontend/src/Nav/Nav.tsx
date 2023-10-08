import React from 'react'
import styles from './Nav.module.css'
import { Link } from "react-router-dom";
import { useState } from 'react'
import walletIcon from './assets/walletIcon.png'; // Import the image
import Vector from './assets/Vector.png'; // Import the image
import burger from './assets/burger.png'; // Import the image

function Nav() {
    // const [ fix,setFix ] = useState(false)
    // function setFixed(){
    //   if (window.scrollY >= 392) {
    //     setFix(true)
    //   }else{
    //     setFix(false)
    //   }
    // }
    // window.addEventListener("scroll",setFixed)
  return (
    <nav id='navbar'>
        <div >
          <Link to='/home' id={styles.burger} >
            <img  src={burger} alt="" />
          </Link>
        </div>
      
        <div id={styles.walletNprofile}>
            <div id={styles.wallet}>
                <Link to='/wallet' id={styles.wallbox} >
                    <h1 id={styles.box} >$1000</h1>
                    <img src={walletIcon} id={styles.box} />
                </Link>
            </div >
            <div id={styles.profile}>
                <a href='#' >Jenny Wilson</a>
                <img src={Vector} />
            </div>
        </div>
    </nav>
  )
}

export default Nav