import React from 'react';

import { goggle, slack, atlassian, dropbox, shopify } from './imports';

import './brand.css';


const Brand = () => {
  return (
    <div classname='gpt3__brand section__padding'>

      <div>
        <img src={goggle} alt="google" />
      </div>

      <div>
        <img src={slack} alt="slack" />
      </div>

      <div>
        <img src={atlassian} alt="atlassian" />
      </div>

      <div>
        <img src={dropbox} alt="dropbox" />
      </div>

      <div>
        <img src={shopify} alt="shopify" />
      </div>

    </div>
  )
}

export default Brand