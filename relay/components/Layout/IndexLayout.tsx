import Head from 'next/head'
import React, { ReactNode } from 'react';
import Nav from '../Nav/nav'
import ProjectList from '../ProjectsList/ProjectsList'
import ActivityFeed from '../ActivityFeed/ActivityFeed';
import ProfileSidebar from '../ProfileSidebar/ProfileSIdebar';

type Props = {
    children?: ReactNode
    title?: string
}


const IndexLayout = ({ children, title = 'Valist' }: Props) => {

    return(
        <div>
            <Head>
                <title>{title}</title>
                <meta charSet="utf-8" />
                <meta name="viewport" content="initial-scale=1.0, width=device-width" />
            </Head>
            <div>
                <div className="fixed top-0 left-0 w-1/2 h-full bg-white"/>
                <div className="fixed top-0 right-0 w-1/2 h-full bg-gray-50"/>
                <div className="relative min-h-screen flex flex-col">
                    <Nav />
                    <div className="flex-grow w-full max-w-7xl mx-auto xl:px-8 lg:flex">
                        <ProfileSidebar />
                        <div className="flex-1 min-w-0 bg-white xl:flex">
                            <ProjectList />
                            <ActivityFeed />
                        </div>
                    </div>
                </div>
            </div>
            {children}
        </div>
    )
}

export default IndexLayout
