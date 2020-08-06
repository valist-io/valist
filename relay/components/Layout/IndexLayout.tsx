import Head from 'next/head'
import React, { useEffect, ReactNode } from 'react';
import Button from '@material-ui/core/Button';
import Link from 'next/link'

type Props = {
    children?: ReactNode
    title?: string
}

const IndexLayout = ({ children, title = 'Valist' }: Props) => {

    useEffect(() => {
        window.onscroll = function() {stickNav()};
        let navbar = document.getElementById("navbar");
        // @ts-ignore
        let sticky = navbar.offsetTop;

        function stickNav() {
            if (window.pageYOffset >= sticky) {
                // @ts-ignore
                navbar.classList.add("sticky")
            } else {
                // @ts-ignore
                navbar.classList.remove("sticky");
            }
        }
    }, []);

    return(
        <div>
        <Head>
            <title>{title}</title>
            <meta charSet="utf-8" />
            <meta name="viewport" content="initial-scale=1.0, width=device-width" />
            <link href="https://fonts.googleapis.com/css2?family=Noto+Sans:wght@400;700&display=swap" rel="stylesheet" />
        </Head>
        <div id="intro">
            <h1>Bring your packages and firmware bundles to the decentralized web!</h1>
            <p>A decentralized binary data notary and global repository, allowing developers and organizations to register public credentials to securely publish and distribute software/firmware/arbitrary data to users. Powered by Ethereum, IPFS, and Filecoin. </p>
            <div id="intro-button">
            <Link href="/org/create">
                <a>
                <Button variant="contained" color="primary">Create an Organization</Button>
                </a>
            </Link>

            </div>
        </div>
        <nav id="navbar">
            <div>Organizations</div>
            <div>Repositories</div>
            <div>Newest Releases</div>
        </nav>
        <div id="valist-content">
            {children}
        </div>
        <footer>
            <hr />
            <span>I'm here to stay (Footer)</span>
        </footer>
    </div>
    )
}

export default IndexLayout
