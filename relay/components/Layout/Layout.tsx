import React, { ReactNode } from 'react'
import Link from 'next/link'
import Head from 'next/head'
import NavBar from '../NavBar/NavBar'

type Props = {
  children?: ReactNode
  title?: string
}

const Layout = ({ children, title = 'Valist' }: Props) => (

  <div>
    <Head>
      <title>{title}</title>
      <meta charSet="utf-8" />
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
    </Head>
    <header>
      <NavBar />
    </header>
    <div id="valist-content">
      {children}
    </div>
    <footer>
        <hr />
        <span>I'm here to stay (Footer)</span>
    </footer>
  </div>
)

export default Layout
