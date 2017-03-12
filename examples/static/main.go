package main

import(
    "github.com/alexbeltran/goChartjs"
    "log"
    "os"
    "io"
    "html/template"
)

type Chart struct{
    goChartjs.Chart
}

func (c *Chart)Get()([]goChartjs.Dataset, error){
    bg := goChartjs.Color{255, 0, 0, 0.4,}
    lineTension := 0;
    return []goChartjs.Dataset{goChartjs.Dataset{
            Label:"Pizza",
            LineTension:&lineTension,
            Data: []goChartjs.Point{
            goChartjs.Point{1318781876406, 0.0},
            goChartjs.Point{1318781976506, 10.0},
            goChartjs.Point{1318792076606, -1.0}},
            BackgroundColor: bg.String(),
        },
    }, nil;
}

func chartHandler(w io.Writer)error{
    tmpl := template.Must(template.ParseFiles("test.html"))
    c := Chart{goChartjs.Chart{
        Name:"pizza",
        ChartType: "line",
        Options: &goChartjs.Options{
            Scales:goChartjs.Scales{
                    XAxes: []goChartjs.Axes{
                        goChartjs.Axes{
                            Time: &goChartjs.Time{
                                Unit: "hour",
                            },
                            Type:"time",
                        },
                    },
            },
            Responsive: true,
        },
    },
    }
    if d, err := c.Get(); err != nil{
        return err;
    }else{
        c.Data.Datasets = d;
    }

    s, err:= c.Render()
    if err != nil{
        return err;
    }
    err = tmpl.Execute(w, struct{Chart template.HTML}{
        template.HTML(s),
    })

    if err != nil{
        return err;
    }
    
    return nil
}

func main(){
    f, err := os.Create("index.html")
    if err != nil{
        log.Println(err)
    }
    defer f.Close()

    err = chartHandler(f);
    if err != nil{
        log.Println(err)
    }
}
