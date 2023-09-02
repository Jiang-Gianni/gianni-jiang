{% func DocTemplate(title string, description string,
        rMap map[string]struct{
    	    Handlers    map[string]string
            MiddleWares []string}, sortedRoutes []string, functionMap map[string]struct {
		File string
		Line int
	}) %}
# **{%s= title %}**

{%s= description %}


## **Routes**

{%for _, pattern := range sortedRoutes %}
{%code funcMap := rMap[pattern] %}
### **{%s pattern%}**
{%if len(funcMap.MiddleWares) > 0 %}
#### **MIDDLEWARES**
{%endif%}
    {%for _, mw := range funcMap.MiddleWares %}
* [_{%s mw %}_]({%s functionMap[mw].File %}#L{%v functionMap[mw].Line %}){%endfor%}

{%for verb, f := range funcMap.Handlers %}
#### **{%s verb %}**
* [_{%s f %}_]({%s functionMap[f].File %}#L{%v functionMap[f].Line %})
{%endfor%}

<br>
<br>

{%endfor%}

{% endfunc %}