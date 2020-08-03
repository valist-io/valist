import React, { FunctionComponent, useState, useEffect, isValidElement  } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            '& > *': {
                margin: theme.spacing(1),
                width: '45ch',

        },
    },
    }),
);

export const CreateOrganizationForm:FunctionComponent<any> = ({contract}) => {

    console.log(contract)
    
    const [organizationShortName, updateOrganizationShortName] = useState("")
    const [organizationFullName, updateOrganizationFullName] = useState("")
    const [organizationDescription, updateOrganizationDescription] = useState("")
    
    const createOrganizationData = () => {

        const meta = {
            Name: organizationFullName,
            Description: organizationDescription 
        };

        alert(organizationShortName)

        contract.methods.createOrganization(organizationShortName, JSON.stringify(meta)).call()
    };

    const classes = useStyles();

    return (
        <form className={classes.root} noValidate autoComplete="off">
            <TextField onChange={(e) => updateOrganizationShortName(e.target.value)} id="outlined-basic" label="Organization Name" variant="outlined" />
            <br></br>
            <TextField onChange={(e) => updateOrganizationFullName(e.target.value)} id="outlined-basic" label="Organization Metadata" variant="outlined" />
            <br></br>
            <TextField onChange={(e) => updateOrganizationDescription(e.target.value)} id="outlined-basic" label="Organization Name" variant="outlined" />
            <br></br>
            <Button onClick={createOrganizationData} value="Submit" variant="contained" color="primary">Create</Button>
        </form>
    );

}