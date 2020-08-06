import React, { useState, useEffect } from 'react';
import IndexLayout from '../components/Layout/IndexLayout'
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';

function MapLsit() {
  const numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9]
  const listItems = numbers.map((number) =>
    <div className="list">
      <div className="list-item">
          <Card>
                  <CardContent>
                    <Typography color="textSecondary" gutterBottom>
                      Organization Name
                    </Typography>
                    <Typography color="textSecondary">
                      {number}
                    </Typography>
                    <Typography variant="body2" component="p">
                      Organization Description
                    </Typography>
                  </CardContent>
                  <CardActions>
                    <Button size="small">Learn More</Button>
                  </CardActions>
                </Card>
      </div>
    </div>
  );

  return listItems
}


export const IndexPage = () => {
  return (
    <IndexLayout title="valist.io">
      <div>
          {MapLsit()}
      </div>
      
    </IndexLayout>
  );
}

export default IndexPage